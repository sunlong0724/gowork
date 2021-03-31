package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main(){
	ch := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(2)
	go func(){
		defer wg.Done()
		defer close(ch)
		for i := 0; i < 4; i++ {
			ch <-rand.Int()
		}
	}()

	go func(){
		defer wg.Done()
		for i := range ch {
			fmt.Println(i)
		}
	}()

	wg.Wait()
	fmt.Println("main exited.")
}