package main

import (
	"fmt"
	"time"
)

func main() {
	message := make(chan int, 10)
	done := make(chan bool)

	defer close(message)

	go func() {
		ticker := time.NewTicker(time.Second)
		for range ticker.C {
			select {
			case <-done:
				fmt.Println("child process interrupt...")
				return
			default:
				fmt.Printf("receive message: %d \n", <-message)
			}
		}
	}()

	for i := 0; i < 10; i++ {
		message <- i
	}
	time.Sleep(5 * time.Second)
	close(done)
	time.Sleep(1 * time.Second)
	fmt.Println("main process exit!")
}
