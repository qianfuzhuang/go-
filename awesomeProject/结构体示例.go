package main

import "fmt"

type person struct {
	Name string
	Age  int
}

func main() {
	stu := person{"xia", 23}
	fmt.Println(stu)
	//指针类型
	stu1 := &person{
		Name: "xue",
		Age:  23,
	}
	fmt.Println(*stu1)
}
