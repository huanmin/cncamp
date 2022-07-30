package b

import (
	"fmt"
	"module1/lifeCycle/a"
)

func init() {
	fmt.Println("b init")
	fmt.Println(a.Sunday)
}
