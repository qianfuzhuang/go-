package main

import (
	"fmt"
	"sync"
	"time"
)

var lock1 sync.RWMutex
var wg1 sync.WaitGroup

func read() {
	defer wg1.Done()
	lock1.RLock()
	fmt.Println("读者读数据")
	time.Sleep(time.Second)
	fmt.Println("读者读取完成")
	lock1.RUnlock()
}
func write() {
	defer wg1.Done()
	lock1.Lock()
	fmt.Println("写者写数据")
	time.Sleep(time.Second * 5)
	fmt.Println("写者写数据完成")
	lock1.Unlock()
}
func main() {
	wg1.Add(6)
	for i := 0; i < 5; i++ {
		go read()
	}
	go write()
	wg1.Wait()
}
