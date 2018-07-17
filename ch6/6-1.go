package main

import "fmt"

func main() {
	a := [...]int{4: 1}
	b := [5]int{4: 1}
	c := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	var p *[5]int = &b
	fmt.Println(p)

	fmt.Println(a == b, a != b)
	fmt.Println(c)
	k := new([10]int)
	k[1] = 10
	fmt.Println(k)
}
