package main

import (
	"fmt"
	"time"
)

func main() {
	intcha := make(chan int, 5)
	go func() {
		time.Sleep(time.Second * 1)
		intcha <- 10
		fmt.Println("int管道传入")
	}()
	strcha := make(chan string, 2)
	go func() {
		time.Sleep(time.Second * 0)
		strcha <- "ijnhgg"
		fmt.Println("str管道传入")
	}()
	select {
	case v := <-intcha:
		fmt.Println(v)
	case v := <-strcha:
		fmt.Println(v)
	default:
		fmt.Println("防止被阻塞")

	}
}
