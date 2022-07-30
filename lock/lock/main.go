package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	go rlock()
	go wlock()
	go lock()
	time.Sleep(5 * time.Second)
}

func lock() {
	lock := sync.Mutex{}
	for i := 0; i < 3; i++ {
		lock.Lock()
		//错误的，释放不了锁，报错
		defer lock.Unlock()
		fmt.Println("lock: ", i)
	}
}

func wlock() {
	lock := sync.RWMutex{}
	for i := 0; i < 3; i++ {
		lock.Lock()
		//错误的，释放不了锁，报错
		defer lock.Unlock()
		fmt.Println("wlock: ", i)
	}
}

func rlock() {
	lock := sync.RWMutex{}
	for i := 0; i < 3; i++ {
		lock.RLock()
		//因为读锁不互斥，循环三次后可以释放
		defer lock.RUnlock()
		fmt.Println("rlock: ", i)
	}
}
