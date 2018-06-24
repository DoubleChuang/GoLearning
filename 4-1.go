package main

import (
	"fmt"
)

const (
	a = "123"
	b = len(a)
	c
	e, f = 1, "2"
	g, h
)

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)

	fmt.Println(Saturday)
}
