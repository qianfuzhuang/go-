package main

import "fmt"

func add(a, b int) int {
	sum := a + b
	return sum
}
func sub(a, b int, fu func(int, int) int) int {
	su := fu(a, b)
	return su
}
func main() {
	fmt.Println(sub(1, 2, add))
	sub(1, 2, func(a int, b int) int {
		if b == 0 {
			println("还可以这样")
		} else {
			println("居然还可以这样")
		}
		return 0
	})

}
