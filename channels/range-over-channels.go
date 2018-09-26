package main

import (
	"fmt"
)

func main() {
	queue := make(chan string, 5)
	queue <- "One"
	queue <- "Two"
	queue <- "Three"
	queue <- "Four"
	queue <- "Five"

	close(queue)

	for q := range queue {
		fmt.Println(q)
	}
}
