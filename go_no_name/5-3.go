package main

import (
	"fmt"
)

func main() {
	a := 1
	for {
		if a > 3 {
			break
		}
		fmt.Println(a)
		a++
	}
	a = 0
	for a < 3 {
		fmt.Println(a)
		a++
	}
	b := "string"
	for i := 0; i < len(b); i++ {
		fmt.Println(string(b[i]))
	}

	fmt.Println("Done")
}
