package main

import (
	"fmt"
	"sync"
)

type Instance struct {
}

// 实例全局对象
var lazyInstance *Instance
var lock = &sync.Mutex{}

func GetInstance() *Instance {
	// 判断第一次如果为空，则给lazy对象重新赋值。
	if lazyInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if lazyInstance == nil {
			fmt.Println("创建实例")
			lazyInstance = &Instance{}
		}
	}
	return lazyInstance
}

func main() {
	// 模拟并发调用实例对象
	for i := 0; i < 100; i++ {
		go func() {
			GetInstance()
		}()
	}

}