package main

import "fmt"

func deleteItem(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func main() {

	//========= slice  =================
	myArray := [5]int{1, 2, 3, 4, 5}
	mySlice := myArray[1:3]
	fmt.Printf("mySlice %+v\n", mySlice)

	fullSlice := myArray[:]
	remove3rdItem := deleteItem(fullSlice, 2)
	fmt.Printf("remove3rdItem %+v\n", remove3rdItem)

	//mySlice1 := new([]int)
	mySlice2 := make([]int, 10)
	fmt.Printf("mySlice2 type=%T, length=%d, cap=%d\n", mySlice2, len(mySlice2), cap(mySlice2))
	mySlice3 := make([]int, 10, 10)
	fmt.Printf("mySlice3 type=%T, length=%d, cap=%d\n", mySlice3, len(mySlice3), cap(mySlice3))
	//========= map ==============

	myMap := make(map[string]string, 10)
	myMap["a"] = "b"
	value, exists := myMap["a"]
	if exists {
		println(value)
	}

	for k, v := range myMap {
		println(k, v)
	}

	myFuncMap := map[string]func() int{
		"funcA": func() int {
			return 1
		},
	}
	fmt.Println(myFuncMap)
	f := myFuncMap["funcA"]
	fmt.Println(f())
}
