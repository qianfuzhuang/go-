package main

import "fmt"

func breakDemo1() {
BREAKDEMO1:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break BREAKDEMO1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	fmt.Println("...")
}
func main() {
	//breakDemo1()
	//a := [5]int{1, 2, 3, 4, 5}
	//s := a[1:3] // s := a[low:high]
	//fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(a))
	//var a = make([]string, 5, 10)
	//for i := 0; i < 10; i++ {
	//	a = append(a, fmt.Sprintf("%v", i))
	//}
	//fmt.Println(a)
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("after init")
	// 对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "小王子"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "沙河"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
}
