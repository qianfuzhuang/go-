package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Open("d:/goland/test.txt")
	if err != nil {
		fmt.Println("文件打开错误")
	}
	fmt.Printf("file=%v", file)
	err1 := file.Close()
	if err1 == nil {
		fmt.Println("关闭成功")

	} else {
		fmt.Println("关闭失败")
	}
	conte, err2 := ioutil.ReadFile("d:/goland/test.txt")
	if err2 == nil {
		fmt.Printf("内容为：%s", conte)
	}
}
