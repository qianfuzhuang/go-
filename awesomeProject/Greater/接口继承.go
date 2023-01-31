package main

import "fmt"

type Ainter interface {
	a()
}
type Binter interface {
	b()
}
type Cinter interface {
	Ainter
	Binter
	c()
}
type stu struct {
}

func (p stu) a() {
	fmt.Println("A方法")
}
func (p stu) b() {
	fmt.Println("B方法")
}
func (p stu) c() {
	fmt.Println("C方法")
}
func main() {
	var s stu
	var d Cinter = s
	d.a()
	d.b()
	d.c()
}
