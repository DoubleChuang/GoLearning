package main

import (
	"fmt"
	"math"
)

type (
	文本 string
)

func main() {
	var (
		str 文本      = "太酷了"
		f   float64 = 1.1
	)
	simple := 1234.1
	i := int(f)

	fmt.Println(i)
	fmt.Println(simple)
	fmt.Println(str)
	fmt.Println(math.MinInt32)

}
