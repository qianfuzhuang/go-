package main

import (
	"fmt"
	"sync"
	"time"
)

var testmap = make(map[int]int, 10)
var wai sync.WaitGroup
var lock sync.Mutex

func jiecheng(num int) {
	defer wai.Done()
	res := 1
	lock.Lock()
	for i := 1; i <= num; i++ {
		res *= num
	}
	testmap[num] = res
	lock.Unlock()
}
func main() {
	start := time.Now()
	fmt.Println(start)

	for i := 0; i < 15; i++ {
		wai.Add(1)
		go jiecheng(i)
	}
	lock.Lock()
	for index, value := range testmap {
		fmt.Printf("第%d个数为：%d\n", index, value)
	}
	lock.Unlock()
	end := time.Now()
	fmt.Println(end)
	wai.Wait()
	//time.Sleep(time.Second * 5)
}
