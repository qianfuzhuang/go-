package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("d:/goland/test1.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("文件读取错误")
	}
	defer file.Close()
	contex := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		s := fmt.Sprintf("qian付壮%d\n", i)
		contex.WriteString(s)
	}
	contex.Flush()
	fmt.Println("文件写入成功")

}
