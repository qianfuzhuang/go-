package main

import "fmt"

type Aint interface {
	A()
}
type Bint interface {
	B()
}
type lex struct {
}

func (p lex) A() {
	fmt.Println("这是A方法")
}
func (p lex) B() {
	fmt.Println("这是B方法")
}
func bang(p Aint) {
	p.A()
}
func bnagb(p Bint) {
	p.B()
}
func main() {
	c := lex{}
	var d lex
	var s Aint = d
	var o Bint = d
	s.A()
	o.B()

	bnagb(c)
	bang(c)
	c.B()
}
