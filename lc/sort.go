package main

import (
	"fmt"
	"reflect"
)

func bubbleSort(s []int){
	for i := 0; i < len(s)-1; i++{
		for j := 1; j < len(s)-i; j++{
			if s[j-1] > s[j]{
				s[j-1],s[j] = s[j], s[j-1]
			}
		}
	}
}

func selectSort(s []int){
	var minI int
	for i := 0; i < len(s)-1; i++{
		minI = i
		for j := i+1; j < len(s); j++{
			if s[j] < s[minI]{
				minI = j
			}
		}
		s[i],s[minI] = s[minI],s[i]
	}
}

func insertSort(s []int){
	for i := 1; i <  len(s); i++ {
		tmp := s[i]
		j := i;
		for ; j >= 1 && s[j-1] > tmp; j--{
			s[j] = s[j-1]
		}
		s[j] = tmp
	}
}

func shellSort(s []int){
	for gap := len(s)/2; gap > 0; gap = gap/2{
		for i := gap; i < len(s); i++{
			tmp := s[i]
			j := i
			for ; j >= gap && s[j-gap] > tmp; j = j-gap {
				s[j] = s[j-gap]
			}
			s[j] = tmp
		}
	}
}

func quickSort(s []int){
	doQuickSort(s, 0, len(s)-1)
}
func doQuickSort(s []int, si,ei int){
	if si >= ei {
		return
	}
	mi := si + (ei-si) >> 1
	s[mi],s[ei] = s[ei],s[mi]
	cur, i := si,si
	for i < ei {
		if s[i] < s[ei] {
			s[i],s[cur] = s[cur],s[i]
			cur++
		}
		i++
	}
	s[cur],s[ei] = s[ei],s[cur]

	doQuickSort(s, si, cur-1)
	doQuickSort(s, cur+1, ei)
}


func mergeSort(s *[]int){
	*s = doMergeSort(*s, 0, len(*s)-1)
}

func doMergeSort(s []int, si, ei int) []int{
	if si > ei {
		return []int{}
	}
	if si == ei {
		return []int{s[si]}
	}
	mi := si + (ei-si)>>1
	ls := doMergeSort(s, si, mi)
	rs := doMergeSort(s, mi+1, ei)
	var ret []int
	i, j := 0,0
	for ; i < len(ls) && j < len(rs);{
		if ls[i] < rs[j]{
			ret = append(ret, ls[i])
			i++
		}else{
			ret = append(ret, rs[j])
			j++
		}
	}
	if i < len(ls){
		ret = append(ret, ls[i:]...)
	}
	if j < len(rs){
		ret = append(ret, rs[j:]...)
	}
	return ret
}

func heapSort(s []int){
	//BUILD MAX HEAP
	for i := len(s) >> 1 -1; i >=0; i--{
		percDown(s, i, len(s))
	}

	for i := len(s)-1; i > 0; i--{
		s[i],s[0] = s[0],s[i]
		percDown(s,0, i)
	}
}
func percDown(s []int, i, l int){
	tmp := s[i]
	for child := 2*i+1; 2*i+1 < l ; i = child{
		child = 2*i+1
		if child + 1 != l && s[child] < s[child+1] {
			child++
		}
		if tmp < s[child]{
			s[i] = s[child]
		}else{
			break
		}
	}
	s[i] = tmp
}

var a,b int

func f() {
	a = 1
	b = 2
}
func g() {
	println(b)
	println(a)
}

func hello(){
	fmt.Println("hello world.")
}
func doCall(callee interface{}){
	ee := reflect.ValueOf( callee)
	fmt.Println(ee.Kind() == reflect.Func)
	ee.Call(nil)
}

type MyType struct{
	i int
	name string
}

func (this *MyType)SetI(i int){
	this.i = i
}
func (this *MyType) SetName(name string){
	this.name = name
}
func (this *MyType) ToString() string{
	return fmt.Sprintf("%p:name:%s,i:%d",this, this.name, this.i)
}
/*
func (this MyType) ToString() string{
	return fmt.Sprintf("%p:name:%s,i:%d",this, this.name, this.i)
}
 */

func main(){
	//var s = []int{2,1,3,2,4,6,5,11,21,22,13,23,33,43,5,4,3,6,8,99}
	//bubbleSort(s)
	//selectSort(s)
	//insertSort(s)
	//shellSort(s)
	//quickSort(s)
	//mergeSort(&s)
	//heapSort(s)
	//var s []int
	//println(s)
	//fmt.Println(len(s))
	//fmt.Println(s)
	//emptyStru := struct {}{}
	//emptyStru2 := struct {}{}
	//println(&emptyStru, &emptyStru2)
	//go f()
	//g()
	//doCall( hello)

	//var mt = MyType{}
	//fmt.Println(mt.ToString())
	//
	//v1 := reflect.ValueOf(&mt).Elem()
	//fmt.Println(v1.MethodByName("ToString").Call(nil)[0])

	//v := reflect.ValueOf(mt)
	//v.MethodByName("SetI").Call([]reflect.Value{ reflect.ValueOf(12)})
	//fmt.Println(v.MethodByName("ToString").Call(nil)[0]	)

}
