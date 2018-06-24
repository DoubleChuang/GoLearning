package main

import (
	"fmt"
)

func main() {
	fmt.Println(^2)
	fmt.Println(1 ^ 2)
	fmt.Println(!true)
	fmt.Println(1 << 8 << 8 >> 8)
	fmt.Println(6 & 11)
	fmt.Println(6 | 11)
	fmt.Println(6 ^ 11)
	fmt.Println(6 &^ 11)
}
