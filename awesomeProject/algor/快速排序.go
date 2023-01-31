package main

import "fmt"

func quicksort(list []int) []int {
	//递归出口
	if len(list) < 2 {
		return list
	} else {
		//将第一个元素设置为基准元素
		pivot := 0
		var item = []int{}
		var item2 = []int{}
		//遍历数组
		for i := 1; i < len(list); i++ {
			//将数组以基准元素划分为两部分
			if list[i] < list[pivot] {
				item = append(item, list[i])
			} else {
				item2 = append(item2, list[i])
			}
		}
		//分别递归划分的两部分再依次划分
		less := append(quicksort(item), list[pivot])
		greater := quicksort(item2)
		return append(less, greater...)
	}

}
func main() {
	fmt.Println(quicksort([]int{11, 22, 1, 5, 99, 6, 5, 77}))
}
