package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(time.Second)
	select {
	case <-timer.C:
		fmt.Println("timeout waiting from channel ch")
	}
}
