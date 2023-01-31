package main

import (
	"fmt"
	"sync"
)

var w sync.WaitGroup

func bianli() {
	defer w.Done()
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
func div() {
	defer w.Done()
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("div函数异常错误")
		}
	}()
	num := 10
	num1 := 0
	num = num / num1
	fmt.Println("ssss", num)
}
func main() {
	w.Add(2)
	go bianli()
	go div()
	w.Wait()
}
