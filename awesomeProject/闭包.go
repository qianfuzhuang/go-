package main

import "fmt"

func main() {
	t1 := test()
	fmt.Println(t1)
	println(t1())
	println(t1())
	println(t1())
	println(t1())
	println(t1())
	t2 := t1()
	println(t2)
	println(t2)
	println(t2)
	t3 := t2
	println(t3)
	println(t1())
	println(t1())
}

func test() func() int {
	a := 0
	gh := func() int {
		a++
		return a
	}
	return gh

}
