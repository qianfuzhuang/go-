package main

import "fmt"

type Peo struct {
	Name string
	Age  int
}

func (s *Peo) String() string {
	stu := fmt.Sprintf("Name=%s\tAge=%d", s.Name, s.Age)
	return stu
}
func main() {
	info := Peo{Name: "qqqqq", Age: 66}
	fmt.Println(&info)
}
