package main

import (
	"fmt"
)

func zeroval(val int) {
	val = 0
}
func zeroptr(val *int) {
	*val = 0
}
func main() {
	i := 1
	fmt.Println("init i:", i)

	zeroval(i)
	fmt.Println("zeroval i:", i)

	zeroptr(&i)
	fmt.Println("zeroptr i:", i)

	fmt.Println("ptr i:", &i)
}
