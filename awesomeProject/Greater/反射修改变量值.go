package main

import (
	"fmt"
	"reflect"
)

type dd struct {
	Sex string
}

func main() {
	//var d int = 50
	//teseref(&d)
	//fmt.Println(d)
	ddd := dd{
		Sex: "女",
	}
	teseref(&ddd)
}
func teseref(i interface{}) {
	ty := reflect.ValueOf(i)
	n := ty.Elem().NumField()
	fmt.Println(n)
	//修改字段值
	ty.Elem().Field(0).SetString("nan")
	//ty.Elem().SetInt(9999)

}
