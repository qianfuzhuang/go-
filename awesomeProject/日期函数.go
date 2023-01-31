package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("日期：%v\n", time.Now())
	fmt.Printf("年：%v\n", time.Now().Year())
	fmt.Printf("月：%v\n", int(time.Now().Month()))
	fmt.Printf("日：%v\n", time.Now().Day())
	fmt.Printf("时：%v\n", time.Now().Hour())
	fmt.Printf("分：%v\n", time.Now().Minute())
	fmt.Printf("星期：%v\n", time.Now().Weekday())
	//fmt.Printf("年：%v\n", time.Now().Year())
	fmt.Println(now.Format("2006/01/02 15:04:05"))

}
