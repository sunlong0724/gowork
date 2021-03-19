package main

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
