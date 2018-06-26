package main

import (
	"fmt"
)

func main() {
	a := 0
	if a, b := 1, 2; a > 0 {
		fmt.Println(a, b)
	}
	fmt.Println(a)
}
