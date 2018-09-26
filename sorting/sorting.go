package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("String:", strs)

	ints := make([]int, 5)
	for i := 0; i < 5; i++ {
		ints[i] = rand.Intn(100)
	}
	sort.Ints(ints)
	fmt.Println("Ints:", ints)
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted:", s)

	floats := make([]float64, 5)
	for i := 0; i < 5; i++ {
		floats[i] = rand.Float64()
	}
	fmt.Println("Floats:", floats)
	sort.Float64s(floats)
	fmt.Println("Floats:", floats)

}
