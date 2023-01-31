package main

import "fmt"

func selectsort(list []int) []int {

	for i := 0; i < len(list); i++ {
		flag := i
		for j := i + 1; j < len(list); j++ {
			if list[flag] > list[j] {
				flag = j
			}
		}
		list[i], list[flag] = list[flag], list[i]
	}
	//fmt.Println(list)
	return list
}
func main() {
	arr := []int{1, 5, 3, 888, 6, 10, 8, 9, 88}
	fmt.Println(selectsort(arr))
}
