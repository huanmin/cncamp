package main

import "fmt"

func main() {
	myArr := []string{"I", "am", "stupid", "and", "weak"}
	result := append(myArr[:2], append([]string{"smart"}, append(myArr[3:4], "strong")...)...)
	fmt.Printf("%+v, len=%d, cap=%d\n", result, len(result), cap(result))

	result1 := answer1(myArr)
	fmt.Printf("%+v, len=%d, cap=%d\n", result1, len(result1), cap(result1))

	result2 := answer2(myArr)
	fmt.Printf("%+v, len=%d, cap=%d\n", result2, len(result2), cap(result2))

}

func answer1(myArr []string) []string {
	result := make([]string, 0)
	for _, value := range myArr {
		if value == "stupid" {
			result = append(result, "smart")
		} else if value == "weak" {
			result = append(result, "strong")
		} else {
			result = append(result, value)
		}
	}
	return result
}

func answer2(myArr []string) []string {
	result := make([]string, 5)
	for index, value := range myArr {
		if index == 2 {
			result[index] = "smart"
		} else if index == 4 {
			result[index] = "strong"
		} else {
			result[index] = value
		}
	}
	return result
}
