package main

import "fmt"

func main() {
	var n Info
	n.Name = "qianffff"

	(&n).people()
	fmt.Println(n.Name)
}

// 传入地址，引用传递
func (q *Info) people() {
	(*q).Name = "qian"
	fmt.Println(q.Name)
}

//结构体传递
//func (q Info) people() {
//	q.Name = "qianfuz"
//	fmt.Println(q)
//}

type Info struct {
	Name string
}
