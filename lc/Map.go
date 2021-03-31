package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
实现阻塞读且并发安全的map
GO里面MAP如何实现key不存在 get操作等待 直到key存在或者超时，保证并发安全，且需要实现以下接口：
type sp interface {
    Out(key string, val interface{})  //存入key /val，如果该key读取的goroutine挂起，则唤醒。此方法不会阻塞，时刻都可以立即执行并返回
    Rd(key string, timeout time.Duration) interface{}  //读取一个key，如果key不存在阻塞，等待key存在或者超时
}
看到阻塞协程第一个想到的就是channel，题目中要求并发安全，那么必须用锁，还要实现多个goroutine读的时候如果值不存在则阻塞，直到写入值，那么每个键值需要有一个阻塞goroutine 的 channel。
*/

type sp interface {
	Out(key string, val interface{})                  //存入key /val，如果该key读取的goroutine挂起，则唤醒。此方法不会阻塞，时刻都可以立即执行并返回
	Rd(key string, timeout time.Duration) interface{} //读取一个key，如果key不存在阻塞，等待key存在或者超时
}

type myMap struct {
	m   map[string]*entry
	rwx sync.RWMutex
}
type entry struct {
	ctx     context.Context
	cancel  context.CancelFunc
	isExist bool
	value   interface{}
}

func (this *myMap) Out(key string, val interface{}) {
	this.rwx.Lock()
	defer this.rwx.Unlock()
	v, ok := this.m[key]
	if !ok {
		this.m[key] = &entry{
			isExist: true,
			value:   val,
		}
		return
	} else {
		if !v.isExist {
			v.isExist = true
			v.value = val
			v.cancel()
		}
	}
	this.m[key].value = val
}

func (this *myMap) Rd(key string, timeout time.Duration) interface{} {
	this.rwx.RLock()
	defer this.rwx.RUnlock()
	v, ok := this.m[key]
	if !ok {
		e := &entry{
			isExist: false,
		}
		e.ctx, e.cancel = context.WithDeadline(context.Background(), time.Now().Add(timeout))
		select {
		case <-e.ctx.Done():
			return e.value
		}
	}
	fmt.Println("here..")
	return v
}

func readM(mm *myMap){
	v := mm.Rd("key1", time.Second* 5)
	fmt.Println(v)
}

func main(){
	mm := myMap{
		m: make(map[string]*entry),
	}
	go readM(&mm)

	go func(){
		mm.Out("key1", "value1")
	}()

	time.Sleep(time.Second* 30)
}