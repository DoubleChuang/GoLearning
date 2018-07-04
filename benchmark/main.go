package main

import (
	"fmt"
	"strconv"
)

func print01(num int) string {
	return strconv.Itoa(num)
}

func print02(num int64) string {
	return strconv.FormatInt(num, 10)
}

func print03(num int) string {
	return fmt.Sprintf("%d", num)
}

func main() {
	fmt.Println(print01(101))
	fmt.Println(print02(101))
	fmt.Println(print03(101))
}
