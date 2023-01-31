package main

import "fmt"

var a []int = []int{1, 2, 3, 4}

func main() {
	fmt.Println(a)
	//二维数组
	var b = [][]int{{1, 2}, {2, 3}}
	var c [2][3]int16
	c[0][1] = 1
	fmt.Println(b)
	//fmt.Println(c)
	var shu [3][3]int = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	//二维数组遍历
	for i := 0; i < len(shu); i++ {
		for j := 0; j < len(shu); j++ {
			println(shu[i][j])

		}
		println()
	}
	for index, valu := range shu {
		for k, v := range valu {
			fmt.Printf("shu[%d][%d]=%d", index, k, v)

		}
	}
	shuz := make([]int, 5)
	shuz[0] = 1
	fmt.Println(shuz)

}
