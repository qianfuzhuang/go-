package main

import "fmt"

func main() {
	var a bool = true
	var b bool = false
	if a && b {
		fmt.Println("a>b")
	} else {
		fmt.Println(!a)
	}
}
