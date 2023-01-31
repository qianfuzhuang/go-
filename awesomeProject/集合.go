package main

import "fmt"

func main() {
	//var a map[int]string

	a := make(map[int]string)
	a[1] = "qqqq"
	a[2] = "wwww"
	a[999] = "eeee"
	fmt.Println(a)
	b := map[int]string{
		111: "qwe",
		123: "qqqq",
	}
	b[521] = "ooooo"
	fmt.Println(b)
	//删除
	delete(b, 111)
	fmt.Println(b)
	value, flag := b[123]
	fmt.Println(value)
	fmt.Println(flag)
	for i, s := range b {
		println(i, s)
	}
	println("--------------------------------")
	c := make(map[string]map[int]string)
	c["qqq"] = map[int]string{
		1: "oooo",
		2: "pppp",
		3: "lili",
	}
	c["www"] = make(map[int]string)
	c["qqq"][6] = "qian"
	c["www"][1] = "test"
	c["www"][2] = "test2"
	fmt.Println(c)
}
