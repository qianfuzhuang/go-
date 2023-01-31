package main

import "fmt"

// 定义接口
type Sayhell interface {
	sayhell()
}

// 接口结构体
type Chines struct {
	name string
}

func (person Chines) sayhell() {
	fmt.Println("你好")
}

func (person Chines) huimian() {
	fmt.Println("烩面")
}

type America struct {
	name string
}

func (person America) sayhell() {
	fmt.Println("Hi")
}
func greete(p Sayhell) {
	p.sayhell()
	/*
		//断言1
		var chi Chines = p.(Chines)
		chi.huimian()
	*/
	//断言2
	chi, flag := p.(Chines)
	if flag {
		chi.huimian()
	} else {
		fmt.Println("有一个美国人不吃烩面")
	}
	//	写法三
	switch p.(type) {
	case Chines:
		ch := p.(Chines)
		ch.huimian()
	}

}
func main() {
	var d Chines
	var a America
	greete(a)
	var s Sayhell = d
	s.sayhell()
	greete(d)
}
