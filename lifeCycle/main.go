package main

import (
	"fmt"
	"learn_go/lifeCycle/a"
	_ "learn_go/lifeCycle/b"
)

func main() {
	fmt.Println(a.Sunday)
}

func add(a, b int) (x, y int) {
	x = a
	y = b
	return x, y
}
