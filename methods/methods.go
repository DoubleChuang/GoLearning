package main

import (
	"fmt"
)

type rect struct {
	width  int
	height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r *rect) perim() int {
	return 2 * (r.width + r.height)
}
func main() {
	r := rect{640, 480}
	fmt.Println("area:", r.area())
	fmt.Println("area:", r.perim())

	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("area:", rp.perim())
}
