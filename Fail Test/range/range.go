package main

import "fmt"

type student struct {
	Name string
	Age  int
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	// wrong way

	//在Go的for range循环中，Go使用copy by value 而不是元素本身
	//當用&取得元素的地址時，其實只是取到copy位置的地址
	//而當跑完最後一次時會留在最後一個的值
	for _, stu := range stus {
		m[stu.Name] = &stu
	}
	for i, v := range m {
		fmt.Println(i, v)
	}

	// right way
	k := make(map[string]*student)
	for i, stu := range stus {
		k[stu.Name] = &stus[i]
	}

	for i, v := range k {
		fmt.Println(i, v)
	}

}
