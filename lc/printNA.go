package main

import (
	"fmt"
	"sync"
)

/*
交替打印数字和字母
使用两个 goroutine 交替打印序列，一个 goroutine 打印数字， 另外一个 goroutine 打印字母， 最终效果如下：
12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
 */
func main(){

	chN := make(chan struct{})
	chA := make(chan struct{})

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func(){
		for i := 1;true;{
			<- chN
			fmt.Printf("%d%d",i,i+1)
			i += 2
			if i >= 29 {
				//fmt.Printf("\n###%d",i)
				//time.Sleep(time.Second)
				wg.Done()
				return
			}
			chA <- struct{}{}
		}
	}()

	go func(){
		for i := 'A';true;{
			<- chA
			fmt.Printf("%c%c", i,i+1)
			i+=2
			chN <- struct{}{}
			if i >= 'Z' {
				//fmt.Printf("\n@@@%d",i)
				//time.Sleep(time.Second)
				wg.Done()
				return
			}
		}
	}()

	chN <- struct{}{}
	wg.Wait()
	close(chN)
	close(chA)
	fmt.Println()
	fmt.Println("main exited...")
}
