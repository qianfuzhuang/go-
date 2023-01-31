package main

import "fmt"

var strch chan string

type dog struct {
	Name  string
	color string
}

func main() {
	//管道操作
	/*
		strch = make(chan string, 5)
		strch <- "111"
		strch <- "ppp"
		fmt.Printf("管道长度为：%d，容量为：%d，第一个数据为：%v\n", len(strch), cap(strch), <-strch)
		mapchan := make(chan map[int]string, 6)
		map1 := make(map[int]string, 5)
		map1[0] = "ppp"
		map1[1] = "kkkkkk"
		map2 := make(map[int]string, 3)
		map2[0] = "9999"
		mapchan <- map1
		mapchan <- map2
		fmt.Printf("管道长度为：%d，容量为：%d，数据为：%v%v", len(mapchan), cap(mapchan), <-mapchan, <-mapchan)
	*/
	allam := make(chan interface{}, 5)
	allam <- dog{
		Name:  "zz",
		color: "yellow",
	}
	allam <- 5
	allam <- "poiu"
	dog1 := <-allam
	//类型断言
	fmt.Printf("dog:%s", dog1.(dog).color)

}
