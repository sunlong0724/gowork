package main

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	cap int
	m   map[string]*list.Element
	l   *list.List
}
type ListNode struct{
	key string
	v   interface{}
}

func (this *LRUCache)init(cap int){
	this.cap = cap
	this.m = make(map[string]*list.Element, 0)
	this.l = list.New()
}

func (this *LRUCache)get(key string) interface{}{
	val, ok := this.m[key]
	if !ok {
		return nil
	}
	this.l.MoveToFront(val)
	return val.Value.(*ListNode).v
}

func (this *LRUCache)put(key string, v interface{}){
	if _, ok := this.m[key]; !ok {
		if len(this.m) == this.cap {
			delete(this.m, this.l.Back().Value.(*ListNode).key)
			this.l.Remove( this.l.Back())
		}
		this.l.PushBack(&ListNode{key,v})
		this.m[key]=  this.l.Back()
	}
	this.l.MoveBefore(this.l.Back(), this.l.Front())
}

func (this* LRUCache)show(){
	e := this.l.Front()
	for e != nil {
		fmt.Println(e.Value.(*ListNode).key, e.Value.(*ListNode).v)
		e = e.Next()
	}
	fmt.Println(len(this.m))
}

func main(){
	var lru LRUCache
	lru.init(4)
	lru.put("key1", 1)
	lru.put("key2", 2)
	lru.put("key3", 3)
	lru.put("key4", 4)

	lru.show()
	fmt.Println("########")
	lru.get("key2")
	lru.show()
	fmt.Println("########")
	lru.put("key5", 5)
	lru.show()
	fmt.Println("########")
	lru.get("key2")
	lru.show()
	fmt.Println("########")
}
