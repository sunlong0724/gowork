package main

import "sync/atomic"

var value int32

func SetValue(delta int32) {
	for {
		v := value
		if atomic.CompareAndSwapInt32(&value, v, (v+delta)) {
			break
		}
	}
}
/*
CompareAndSwapInt32 函数不需要循环调用.
 */