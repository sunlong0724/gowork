package main

import (
	"fmt"
)

func Iter() <-chan interface{} {
	ch := make(chan interface{}, 1)
	ch <- 1
	close(ch)
	return ch
}

func main() {

	ch := Iter()
	for {
		select {
		case v, ok := <- ch:
			fmt.Println(v, ok)
			if ok == false {
				return
			}
		}

	}

	//runtime.GOMAXPROCS(1)
	//int_chan := make(chan int, 1)
	//string_chan := make(chan string, 1)
	//int_chan <- 1
	//string_chan <- "hello"
	//select {
	//case value := <-int_chan:
	//	fmt.Println(value)
	//case value := <-string_chan:
	//	panic(value)
	//}
}
