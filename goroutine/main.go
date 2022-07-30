package main

import (
	"fmt"
	"time"
)

func main() {
	//没有缓冲区，两边堵塞
	//ch := make(chan int)
	ch := make(chan int, 10)

	//单向只发送通道
	//var sendOnly chan<- int
	//单向只接收通道
	//var readOnly <-chan int

	go func() {
		fmt.Println("from child thread")
		ch <- 1
		//当使用遍历通道时，必须关闭
		close(ch)
	}()

	//a := <- ch
	//fmt.Println(a)
	for v := range ch {
		fmt.Println("receiving: ", v)
	}

	c := make(chan int)
	go prod(c)
	go consume(c)

	time.Sleep(2 * time.Second)
}

func prod(ch chan<- int) {
	defer close(ch)
	for {
		ch <- 1
	}
}

func consume(ch <-chan int) {
	for {
		<-ch
	}
}
