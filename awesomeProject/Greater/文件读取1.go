package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("d:/goland/test.txt")
	if err != nil {
		fmt.Printf("文件读取失败")
	}
	defer file.Close()
	contex := bufio.NewReader(file)
	for {
		line, err1 := contex.ReadString('\n')
		if err1 == io.EOF {
			fmt.Printf("%s\n", line)
			fmt.Println("读取结束")
			break
		}
		fmt.Printf("%s", line)
	}
}
