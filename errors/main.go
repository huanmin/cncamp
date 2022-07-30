package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type ApiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *ApiError) Error() string {
	errRep, err := json.Marshal(e)
	if err != nil {
		fmt.Println(err)
	}
	return string(errRep)
}

func main() {

	err := fmt.Errorf("this is a error")
	fmt.Println(err)
	errNotFound := errors.New("NotFound")
	fmt.Println(errNotFound)
}
