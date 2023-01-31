package main

import "fmt"

// 两个数的和
func num(a, b int) int {
	c := a + b
	return c
}

// 输出字符串
func prin(mes string) {
	println(mes)
}

// 交换两个字符串位置
func swap(a, b string) (string, string) {
	b, a = a, b
	return a, b
}
func main() {
	//println(swap("qian", "fu"))
	//prin("zhuang")
	//println(num(44, 55))
	getnum("qqqqqqqq", 1, 2, 3, 4, 7, 8, 66, 999, 88888)

}

// 传递多个参数
func getnum(str string, nums ...int) {
	num := 0
	for i := 0; i < len(nums); i++ {
		num += nums[i]
	}
	fmt.Printf("字符串为：%s\t和为：%d", str, num)
}
