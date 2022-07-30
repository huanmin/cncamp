package main

import (
	"fmt"
	"time"
)

func main() {
	// fmt.Println("hello world")

	queue := make(chan int, 10)
	done := make(chan bool)

	go func() {
		i := 0
		ticker := time.NewTicker(2 * time.Second)
		for range ticker.C {
			select {
			case <-done:
				fmt.Println("child process interrupt...")
				return
			default:
				i++
				fmt.Println("向队列里插入：", i)
				queue <- i
			}
		}
	}()

	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for range ticker.C {
			fmt.Println("从队列里取：", <-queue)
		}
	}()

	time.Sleep(10 * time.Second)
	close(done)
	time.Sleep(1 * time.Second)
	fmt.Println("main process exit!")
}
