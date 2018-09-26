package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}
type rect struct {
	width  float64
	height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2 * (r.width + r.height)
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (c circle) perim() float64 {
	return 2 * c.radius * math.Pi
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
func main() {
	r := rect{10, 10}
	c := circle{10}

	measure(r)
	measure(c)
}