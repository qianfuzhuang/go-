package main

import (
	"fmt"
	"time"
)

func test() {
	for i := 0; i < 10; i++ {
		fmt.Printf("这是第%d个输出\n", i)
		time.Sleep(time.Second)
	}
}
func main() {
	go test()
	for j := 0; j < 5; j++ {
		fmt.Printf("学习的第%d天\n", j)
		time.Sleep(time.Second)
	}
}
