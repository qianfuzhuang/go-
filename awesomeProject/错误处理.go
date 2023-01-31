package main

import (
	"errors"
	"fmt"
)

func tes() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	num1 := 5
	num2 := 0
	sum := num1 / num2
	fmt.Println(sum)

}
func diy() (err error) {
	num1 := 5
	num2 := 0
	if num2 == 0 {
		return errors.New("除数不能为0")
	} else {
		result := num1 / num2
		fmt.Println(result)
		return nil
	}
}
func main() {
	tes()
	err := diy()
	if err != nil {
		fmt.Println("自定义错误：", err)
		panic(err)
	}
	println("错误已经跳过")
}
