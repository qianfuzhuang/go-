package main

import "fmt"

type info struct {
	Nane string
	Age  int
	sex  string
}

// 方法四
func main() {
	var stu *info = &(info{"qian", 23, "男"})
	fmt.Println(*stu)
}

// 方法三
//func main() {
//	var stu *info = new(info)
//	stu.Nane = "qian"
//	stu.Age = 23
//	stu.sex = "男"
//	fmt.Println(*stu)
//}

//方法二
//func main() {
//	var stu info = info{"I love zcj", 23, "男"}
//	fmt.Println(stu)
//}
//f方法一
//func main() {
//	var stu info
//	stu.Nane = "qian付"
//	stu.sex = "男"
//	stu.Age = 23
//	fmt.Println(stu)
//
//}
