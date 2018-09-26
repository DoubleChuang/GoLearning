package main

import (
	"fmt"
)

func main() {
	a := 1
	a++
	a--
	var p *int = &a
	fmt.Println(*p)
}
