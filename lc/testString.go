package main

import "fmt"

type People01 struct {
	Name string
}

func (p *People01) String() string {
	return fmt.Sprintf("print: %v", p)
}

func main() {
	p := &People01{}
	p.String()
	fmt.Println(p.String())
}

/*
type Student struct {
	number int
	realname string
	age int
}
func (t *Student) String() string {
	return fmt.Sprintf("学号: %d\n真实姓名: %s\n年龄: %d\n", t.number, t.realname, t.age)
}
func main() {
	stu := &Student{
		number:   1,
		realname: "王小明",
		age:      18,
	}
	//fmt.Println(stu)
	stu.String()
}
 */