package main

import "fmt"

func main() {
	/*
		var i int
		var f float64
		var b bool
		var s string
		fmt.Printf("%v %v %v %q\n", i, f, b, s)
		const leng int = 159
		const wid int = 357
		var area int
		const a, n, c = 1, false, "str"
		area = leng * wid
		fmt.Printf("面积为：%d", area)
		println()
		println(a, b, c)
		const (
			Unknow = "asdf"
			k      = len(Unknow)
			l      = unsafe.Sizeof(Unknow)
		)
		println(Unknow, k, l)
	*/
	const (
		a = 1 << iota
		b = 3 << iota
		c
		d
		e = "ha"
		f
		g = 100
		h
		j = iota
		l
	)
	fmt.Println(a, b, c, d, e, f, g, h, j, l)
}
