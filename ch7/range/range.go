package main

import "fmt"

func main() {
	sm := make([]map[int]string, 5) //slice of map[int]string

	for _, v := range sm {
		v = make(map[int]string, 1)
		v[1] = "OK"
		fmt.Println(v)
	}
	fmt.Println(sm)

	for k := range sm {
		sm[k] = make(map[int]string, 1)
		sm[k][1] = "OK"
		fmt.Println(sm[k])
	}
	fmt.Println(sm)

}
