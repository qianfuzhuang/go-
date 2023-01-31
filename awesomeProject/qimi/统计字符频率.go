package main

import "fmt"

//	func main() {
//		var s = "how do you do"
//		var wordcou = make(map[string]int)
//		str := strings.Split(s, " ")
//		for _, s2 := range str {
//			wordcou[s2]++
//
//		}
//		fmt.Println(wordcou)
//	}
func main() {
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%+v\n", s)
	m["q1mi"] = s
	s = append(s[:1], s[2:]...)
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", m["q1mi"])
}
