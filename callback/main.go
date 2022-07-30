package main

import "fmt"

type Person struct {
	Name string
}

//这是一个方法
func (p *Person) getName() string {
	return p.Name
}

func main() {

	p := Person{Name: "张三"}
	fmt.Println(p.getName())

	func() {
		fmt.Println("这是一个匿名函数")
	}()

	DoOperation(1, increase)
	DoOperation(1, decrease)
}

//以下都是函数
func decrease(a, b int) {
	fmt.Println("increase result is: ", a-b)
}

func increase(a, b int) {
	fmt.Println("increase result is: ", a+b)
}

func DoOperation(y int, f func(int, int)) {
	f(y, 1)
}
