package main

import "fmt"

func f1(a, b int) {
	fmt.Println(a, b)
}

func main() {
	fmt.Printf("类型为：%T\n", f1)
	var f2 func(int, int)
	f2 = f1
	f3 := f1
	f2(1, 2)
	f3(4, 5)
	a := func() {
		println("ceshi")
	}
	a()
	//匿名函数自己调用自己
	func() {
		println("匿名函数")
	}()
	//	匿名函数可以传参
	sum := func(a, b int) int {
		println(a, b)
		println("匿名函数")
		return (a + b)
	}(99, 88)
	println(sum)

}
