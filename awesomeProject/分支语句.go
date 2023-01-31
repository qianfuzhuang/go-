package main

import "fmt"

func main() {
	var sore int
	fmt.Scan(&sore)
	switch sore {
	case 90:
		println("A+")
	case 80:
		println("B")
	case 70:
		println("C")
	default:
		println("不及格")

	}
	switch {
	case true:
		println("1")
		fallthrough
	case false:
		println("2")
	}
}
