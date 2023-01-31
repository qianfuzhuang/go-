package main

import "fmt"

func quicksort(list []int) []int {
	if len(list) < 2 {
		return list
	} else {
		pivot := 0
		var item = []int{}
		var item2 = []int{}
		for i := 1; i < len(list); i++ {
			if list[i] < list[pivot] {
				item = append(item, list[i])
			} else {
				item2 = append(item2, list[i])
			}
		}
		less := append(quicksort(item), list[pivot])
		greater := quicksort(item2)
		return append(less, greater...)
	}

}
func main() {
	fmt.Println(quicksort([]int{11, 22, 1, 5, 99, 6, 5, 77}))
}
