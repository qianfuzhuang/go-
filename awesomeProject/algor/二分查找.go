package main

import "fmt"

func binsearch(arr []int, key int) int {
	low := 0
	high := len(arr)
	for low < high {
		//中间元素位置
		mid := (low + high) / 2
		if key == arr[mid] {
			return mid
		}
		if key < arr[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
func main() {
	arr := []int{1, 2, 3, 5, 6, 7, 8, 9, 88}
	succe := binsearch(arr, 9)
	if succe == -1 {
		println("查找失败")
	} else {
		fmt.Printf("在数组中的位置是：arr[%d]", succe)
	}
}
