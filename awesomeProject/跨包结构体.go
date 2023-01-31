package main

import (
	"awesomeProject/db"
	"fmt"
)

func main() {
	stu := db.Demo{"夏", 56}
	fmt.Println(stu)
	stu1 := db.NewDemo("雪", 23)
	fmt.Println(*stu1)
	bao := db.NewPer("雪")
	bao.SetAge(23)
	fmt.Println(bao.GetAge())
}
