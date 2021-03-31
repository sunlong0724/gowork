package main

import "fmt"

const (
	aa = iota
	bb = iota
)
const (
	name = "menglu"
	c    = iota
	d    = iota
)

func main() {
	fmt.Println(aa)
	fmt.Println(bb)
	fmt.Println(c)
	fmt.Println(d)
}
