package main

import "fmt"

type IF interface {
	getName() string
}

type Human struct {
	firstName, lastName string
}

type Plane struct {
	vendor string
	model  string
}

type Car struct {
	factory, model string
}

func (h *Human) getName() string {
	return h.firstName + "," + h.lastName
}

func (p Plane) getName() string {
	return fmt.Sprintf("vendor: %s, model: %s", p.vendor, p.model)
}

func (c *Car) getName() string {
	return c.factory + "-" + c.model
}

func main() {
	interfaces := []IF{}
	h := new(Human)
	h.firstName = "first"
	h.lastName = "last"
	interfaces = append(interfaces, h)
	//这个不能用，有指针接收器
	//c := Car{factory: "benz", model: "s"}
	c := new(Car)
	c.factory = "benz"
	c.model = "s"
	interfaces = append(interfaces, c)
	//这个没用指针，可以接收
	p := Plane{vendor: "testVendor", model: "testModel"}
	interfaces = append(interfaces, p)
	for _, f := range interfaces {
		fmt.Println(f.getName())
	}
	//p := Plane{}
	//p.vendor = "testVendor"
	//p.model = "testModel"
	//fmt.Println(p.getName())
}
