package main

import (
	"fmt"
	"reflect"
)

type stud struct {
	Name string
	Sex  int
}

func main() {
	num := 100
	info := stud{
		Name: "雪雪",
		Sex:  23,
	}
	reflecttest(info)
	reflecttest(num)
}
func reflecttest(i interface{}) {
	//返回类型
	ty := reflect.TypeOf(i)
	fmt.Println(ty)
	//返回类型数据
	fmt.Printf("类型为：%T\n", ty)
	va := reflect.ValueOf(i)
	fmt.Println(va)
	fmt.Printf("类型为：%T\n", va)
	//此时的数据需要转换才可以使用
	//val := va.Int()
	//fmt.Println(val)
	//	转空接口
	n := va.Interface()
	//类型断言
	num, err := n.(stud)
	if err == true {
		fmt.Printf("学生姓名为：%s,学生年龄是：%d", num.Name, num.Sex)
		return
	}
	//num = num + 60
	//fmt.Println(num)
}
