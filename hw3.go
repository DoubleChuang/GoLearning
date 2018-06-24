package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int = 65
	b := string(a)
	c := strconv.Itoa(65)
	d, _ := strconv.Atoi(c)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

}
