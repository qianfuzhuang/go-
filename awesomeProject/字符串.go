package main

import "fmt"

func main() {
	var str string = "qianfuzhuang"
	for i := 0; i < len(str); i++ {
		fmt.Print(str[i])
		fmt.Printf("%c", str[i])

	}
	println()
	for i2, i3 := range str {
		fmt.Print(i2)
		fmt.Printf("%c\n", i3)

	}
}
