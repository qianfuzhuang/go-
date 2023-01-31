package main

import "fmt"

func isprim(n int) {
	for i := 1; i <= n; i++ {
		var flag bool = true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				flag = false
				continue
			}
		}
		if flag {
			fmt.Printf("%d是素数\n", i)

		}

	}
}
func main() {
	isprim(200)
	fmt.Println("结束")
}
