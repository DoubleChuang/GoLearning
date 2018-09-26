package main

import (
	"fmt"
)

func sum(nums ...int) int {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Printf("%d\n", total)
	return total
}

func main() {
	sum(1, 2, 3)
	sum([]int{1, 2}...)
}
