package main

import "fmt"

var intchannl chan int

func main() {
	intchannl := make(chan int, 103)
	fmt.Printf("类型为：%v\n", intchannl)
	for i := 0; i < 100; i++ {
		intchannl <- i
	}

	intchannl <- 10
	num := 999
	intchannl <- num
	intchannl <- 8654
	num1 := <-intchannl
	num2 := <-intchannl
	close(intchannl)
	for v := range intchannl {
		fmt.Printf("遍历管道元素：%v\n", v)
	}
	num3 := <-intchannl
	fmt.Println(num3)
	fmt.Println(num1, num2)
	fmt.Printf("管道中数据长度为：%v,管道容量为：%v", len(intchannl), cap(intchannl))
}
