package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file1path := "d:/goland/test.txt"
	file2path := "d:/goland/test1.txt"
	red, err := ioutil.ReadFile(file1path)
	if err != nil {
		fmt.Println("读取错误")
	}
	wri := ioutil.WriteFile(file2path, red, 0666)
	if wri != nil {
		fmt.Println("写入失败")
	}
}
