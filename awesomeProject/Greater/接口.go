package main

import "fmt"

// 定义接口
type Sayhello interface {
	sayhello()
}

// 接口结构体
type Chinese struct {
	name string
}

func (person Chinese) sayhello() {
	fmt.Println("你好")
}

type American struct {
	name string
}

func (person American) sayhello() {
	fmt.Println("Hi")
}
func greet(p Sayhello) {
	p.sayhello()
}

// 自定义数据类型
type inter int

func (i inter) sayhello() {
	fmt.Println("hizong+", i)
}

func main() {
	var s inter = 10
	s.sayhello()

	c := Chinese{}
	a := American{}
	var f American
	var g Sayhello = f
	g.sayhello()
	var d Sayhello = c
	d.sayhello()
	greet(c)
	greet(a)
	//多态数组
	var arr [2]Sayhello
	arr[0] = American{"qian"}
	arr[1] = Chinese{"钱"}
	fmt.Println(arr)
}
