package main

import "fmt"

func main() {
	var k map[int]string = map[int]string{}
	m := make(map[int]string)
	m[1] = "OK"

	fmt.Println("k:", k, " m:", m)
	delete(m, 1)
	fmt.Println("k:", k, " m:", m)

	var m1 map[int]map[int]string
	m1 = make(map[int]map[int]string)
	m1[0] = make(map[int]string)
	m1[0][1] = "OK"
	fmt.Println(m1)
	_, ok := m1[2][1]
	if !ok {
		m1[2] = make(map[int]string)
		m1[2][1] = "GOOD"
	}
	a, ok := m1[2][1]
	fmt.Println(a, ok)
}
