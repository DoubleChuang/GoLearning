package main

import (
	"fmt"
)

func main() {
	defer_call()
}

func defer_call() {
	defer func() { fmt.Println("1") }()
	defer func() { fmt.Println("2") }()
	defer func() { fmt.Println("3") }()

	panic("panic")
}
