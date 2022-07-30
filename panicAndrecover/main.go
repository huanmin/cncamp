package main

import "fmt"

func main() {

	err := fmt.Errorf("this is a error")
	fmt.Println(err)
	defer func() {
		fmt.Println("defer func is called")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	panic("a panic is triggered")
}
