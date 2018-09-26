package main

import (
	"fmt"
)

func nine(x int, y int) {
	if x < 10 {
		if y < 10 {
			fmt.Printf("| %2d x %2d = %2d ", x, y, x*y)
			nine(x, y+1)
		} else {
			fmt.Println("|")
			nine(x+1, 1)
		}
	}

}
func main() {
	nine(1, 1)
}
