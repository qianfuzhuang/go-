package main

import (
	"fmt"
	"sync"
)

var sum int
var pv sync.WaitGroup
var lock sync.Mutex

func add() {
	defer pv.Done()
	for i := 0; i < 123456; i++ {
		lock.Lock()
		sum = sum + 1
		lock.Unlock()

	}
}
func sub() {
	defer pv.Done()
	for i := 0; i < 123456; i++ {
		//上锁互斥访问临界区
		lock.Lock()
		sum = sum - 1
		lock.Unlock()
	}
}
func main() {
	pv.Add(2)
	go add()
	go sub()
	pv.Wait()
	fmt.Println(sum)
}
