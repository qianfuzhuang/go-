package main

import "fmt"

const secr int = 12345678

func main() {
	yanz()

}
func yanz() {
	var a int
	var b int
	fmt.Scanln(&a)
	if a == secr {
		fmt.Println("请再次输入密码：")
		fmt.Scanln(&b)
		if b == secr {
			fmt.Println("登录成功！！")
		} else {
			fmt.Println("两次密码不一致！")
		}
	} else {
		fmt.Println("密码错误重新输入")
		yanz()
	}
}
