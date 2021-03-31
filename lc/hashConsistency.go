package main

import (
	"fmt"
	"hash/crc32"
	"sort"
)

type uintSlice []uint32

func (this uintSlice) Len() int {
	return len(this)
}

func (this uintSlice) Swap(i, j int) {
	(this)[i], (this)[j] = (this)[j], (this)[i]
}
func (this uintSlice) Less(i, j int) bool {
	return (this)[i] < (this)[j]
}

type HashConisistency struct {
	nodesMap        map[uint32]string
	nodesForMapSort uintSlice
	nodeReplicas    map[string]uintSlice
	nodes           []string
	replicas        int
}

func (this *HashConisistency) Init(nodes []string, replicas int) {
	this.nodesMap = map[uint32]string{}
	this.nodesForMapSort =  uintSlice{}
	this.nodeReplicas = map[string]uintSlice{}
	this.replicas = replicas
	this.nodes = nodes
	for _, v := range this.nodes{
		this.addNode(v)
	}
	this.sortNodes()
}

func (this *HashConisistency) GetNode(key string) string {
	key_hash := crc32.ChecksumIEEE([]byte(key))
	for _, v := range this.nodesForMapSort {
		if v < key_hash {
			continue
		}
		return this.nodesMap[key_hash]
	}
	return ""
}

func (this *HashConisistency) AddNode(node string) {
	this.addNode(node)
	this.sortNodes()
}

func (this *HashConisistency) RemoveNode(node string) {
	if _,ok := this.nodeReplicas[node];ok {
		discardNodes := this.nodeReplicas[node]
		for _, v := range discardNodes {
			delete(this.nodesMap, v)
			for i, vv := range this.nodesForMapSort {
				if v == vv {
					this.nodesForMapSort = append(this.nodesForMapSort[:i], this.nodesForMapSort[i+1:]...)
				}
			}
		}
	}
}

func (this *HashConisistency) addNode(node string) {

	var replS []uint32
	for i := range this.nodeReplicas {
		rep_node := fmt.Sprintf("%s_%d", node, i)
		node_hash := crc32.ChecksumIEEE([]byte(rep_node))
		this.nodesMap[node_hash] = node
		replS = append(replS, node_hash)

		this.nodesForMapSort = append(this.nodesForMapSort, node_hash)
	}
	this.nodeReplicas[node] = replS
}

func (this *HashConisistency) sortNodes() {
	sort.Sort(this.nodesForMapSort)
}

func main(){
	var memcache_servers = []string{
		"127.0.0.1:7001",
		"127.0.0.1:7002",
		"127.0.0.1:7003",
		"127.0.0.1:7004",
	}
	var h HashConisistency
	h.Init(memcache_servers, 5)

	fmt.Println(h.nodesMap)
	fmt.Println(h.nodesForMapSort)

}