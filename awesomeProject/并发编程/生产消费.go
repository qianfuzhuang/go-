package main

import (
	"fmt"
	"strconv"
	"sync"
)

type product struct {
	Name string
}

// 生产商品放入管道
func Produc(storechan chan<- product, count int) {
	//defer war.Done()

	for {
		produ := product{Name: "商品" + strconv.Itoa(count)}
		storechan <- produ
		count--
		fmt.Println("生产了", produ)
		if count == 0 {
			return
		}
	}
}
func Transport(storchan <-chan product, shopchan chan<- product) {
	//defer war.Done()
	for {
		something := <-storchan
		fmt.Println("运输了商品：", something)
		shopchan <- something
	}

}
func Customer(shopchan chan product, count int, exitchan chan bool) {
	//defer war.Done()
	for {
		custom := <-shopchan
		count--
		fmt.Println("消费了商品：", custom)
		if count == 0 {
			exitchan <- true
			return
		}
	}
}

var war sync.WaitGroup

func main() {
	var storechan chan product = make(chan product, 3000)
	shopchan := make(chan product, 3000)
	exitchan := make(chan bool, 1)
	//war.Add(3)
	go Produc(storechan, 2000)
	go Transport(storechan, shopchan)
	go Customer(shopchan, 200, exitchan)
	if <-exitchan {
		return
	}
	//war.Wait()
	//time.Sleep(time.Second * 5)
}
