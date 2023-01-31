package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var sortma = make(map[string]int, 400)
	for i := 0; i < 200; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(200)
		sortma[key] = value
	}
	var sorte []string
	//soree:=make([]string,0,200)
	for s, _ := range sortma {
		sorte = append(sorte, s)

	}
	sort.Strings(sorte)
	for _, s := range sorte {
		fmt.Printf("%v：%v\n", s, sortma[s])
	}
}
