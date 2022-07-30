package main

import (
	"encoding/json"
	"fmt"
)

type Human struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	h := Human{Name: "张三", Age: 28}
	jsonData, err := json.Marshal(&h)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonData))

	str := "{\"name\":\"张三\",\"age\":28}"
	h1 := Human{}
	err1 := json.Unmarshal([]byte(str), &h1)
	if err1 != nil {
		println(err1)
	}
	fmt.Println(h1)
}
