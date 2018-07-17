package main

import "fmt"

func main() {
	var s1 []int
	fmt.Println(s1)

	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(a)

	s1 = a[5:]
	fmt.Println(s1)

	s2 := make([]int, 3, 10)
	fmt.Println(s2, len(s2), cap(s2))

	b := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k'}
	s_1 := b[3:7]
	s_2 := b[4:]
	fmt.Println(string(s_1), string(s_2))

	b[5] = 'z'
	fmt.Println(string(s_1), string(s_2))
	fmt.Println(len(s_1), cap(s_1))
	fmt.Println(len(s_2), cap(s_2))

	s_3 := make([]int, 3, 5)
	fmt.Printf("%v %p\n", s_3, s_3)
	s_3 = append(s_3, 1, 2, 3)
	fmt.Printf("%v %p\n", s_3, s_3)

	s_4 := []int{1, 2, 3, 4, 5, 6}
	s_5 := []int{7, 8, 9}
	copy(s_5, s_4)
	fmt.Println(s_5)
	copy(s_5, s_4[3:])
	fmt.Println(s_5)
}
