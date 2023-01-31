package main

import "fmt"

func main() {
	//只写管道
	var chann chan<- int
	chann = make(chan int, 6)

	for i := 0; i < 3; i++ {
		chann <- i

	}
	fmt.Printf("管道写入完成")
	//	只读管道
	var rchan <-chan int
	if rchan != nil {
		num := <-rchan
		fmt.Println(num)
	}

}
