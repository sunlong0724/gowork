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


type Peoplee struct{}

func (p *Peoplee) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *Peoplee) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	Peoplee
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func main111() {
	t := Teacher{}
	t.ShowA()
}


//func main(){
//	for i := 0; i < 100; i++{
//		fmt.Println(Selector())
//	}
//}

type people interface {
	Show()
}

type student struct{}

func (this *student)Show(){
}

func live() people{
	var stu *student
	return stu
}

func main(){
	if live() == nil {
		fmt.Println("AAAAAAAAAAAAAA")
	}else{
		fmt.Println("BBBBBBBBBBBBB")
	}
}
