package main

import "fmt"

func main() {
	var a = []int{1, 2, 3}
	var b = make([]int, 5)
	b[2] = 5
	//slice := a[1:3]
	fmt.Println(append(a, b...))
	copy(b, a)
	fmt.Println(b)
}
