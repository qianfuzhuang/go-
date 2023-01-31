package main

import (
	"fmt"
	"sync"
	"time"
)

var intchan chan int
var pvs sync.WaitGroup

// 写管道
func writ(intchan chan int) {
	defer pvs.Done()
	for i := 0; i < 10; i++ {
		intchan <- i
		fmt.Printf("写入的数据为：%d\n", i)
		time.Sleep(time.Second * 2)
	}
	close(intchan)
}

// 读管道
func readd(intchan chan int) {
	defer pvs.Done()
	for i := range intchan {

		fmt.Printf("读取的数据为：%d\n", i)
	}
}

func main() {
	intchan = make(chan int, 50)
	pvs.Add(2)
	go writ(intchan)
	go readd(intchan)
	pvs.Wait()
}
