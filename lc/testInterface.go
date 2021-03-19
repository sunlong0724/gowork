package main

import (
	"fmt"
	"math/rand"
)

type Cat struct{
}
type Duck interface {
	Quack()
}

//func (c Cat)Quack(){
//
//}
func (c *Cat)Quack(){

}

//var d Duck = Cat{}  用指针类型实现接口方法，结构体变量是没有实现接口的。
//var d Duck = &Cat{}

type Server int

var (
	servers  = [10]Server{1,2,3,4,5,6,7,8,9,10}
)

func Selector() Server{
	return servers[ rand.Intn(len(servers))]
}

func main(){
	for i := 0; i < 100; i++{
		fmt.Println(Selector())
	}
}