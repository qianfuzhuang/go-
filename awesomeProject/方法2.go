package main

import "fmt"

type Date string

func (stu *Date) Stud() {
	*stu = "111"
}

func (stu Date) Stude() {
	stu = "5555"
}
func main() {
	var n Date = "ppp"
	n.Stud()
	n.Stude()
	fmt.Println(n)
}
