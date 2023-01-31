package main

import (
	"fmt"
	"reflect"
)

type teach struct {
	Name string
	Age  int
}

func (t teach) Apgread() {
	fmt.Println("升职方法")
}
func (t teach) Bog(a, b int) int {
	return a + b
}
func (t teach) Set(a int, b string) int {
	t.Age = a
	t.Name = b

	fmt.Println("修改字段")
	return a
}
func opeastu(t interface{}) {
	res := reflect.ValueOf(t)
	fmt.Println("反射值：", res)
	nub := res.NumField()
	fmt.Println("结构体内部字段", nub)
	//	遍历每个字段
	for i := 0; i < nub; i++ {
		fmt.Printf("第%d个字段值为：%v\n", i, res.Field(i))
	}
	me := res.NumMethod()
	fmt.Println(me)
	var parm []reflect.Value
	parm = append(parm, reflect.ValueOf(99))
	parm = append(parm, reflect.ValueOf("zcj"))
	sult := res.Method(2).Call(parm)
	fmt.Println(sult[0].Int())
}
func main() {
	opeastu(teach{
		Name: "婧婧",
		Age:  23,
	})
}
