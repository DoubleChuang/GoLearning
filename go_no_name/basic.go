package main

import (
	std "fmt"
	. "os"
)

const PI = 3.14

//name := "gopher" // (x)if global
var (
	name       = "gopher"
	num  int64 = 123

	numPi = PI
)

type (
	newType int
	gopher  struct{}
	golang  interface{}
)

func main() {
	std.Println("Hello world! 你好 世界!")
	std.Fprintf(Stderr, "name:%s num:%d numPi:%f PI:%f\n", name, num, numPi, PI)
	Exit(0)
}
