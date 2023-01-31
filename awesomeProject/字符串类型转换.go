package main

import (
	"fmt"
	"strconv"
	"strings"
)

var str string = "qianfuzhuang"
var chinese = "钱付壮"

func main() {
	println(len(str))
	qiang := []byte(str)
	qiang[0] = 'p'
	fmt.Println(string(qiang))
	for _, i2 := range chinese {
		fmt.Printf("字符串：%c\n", i2)
	}
	str2, _ := strconv.Atoi("lll")
	println(str2)
	//注意汉字的字节
	for i := 0; i < len(chinese); i++ {
		fmt.Printf("%v(%c) ", chinese[i], chinese[i])
	}
	s := []rune(chinese)
	fmt.Printf("%c", s)
	//整型转字符串
	num := strconv.Itoa(9999)
	println(num)
	re := strings.Contains("ilovegolang", "go")
	println(re)
	tet := strings.Count("qianggggjjj", "g")
	println(tet)
	//==返回0，包含返回1 不同返回-1
	ret := strings.Compare("asd", "asd")
	println(ret)
	//不区分大小写比较
	eq := strings.EqualFold("GO", "go")
	println(eq)
	//返回第一次出现的索引值
	fmt.Println(strings.Index("ilovego", "go"))
	//字符串替换
	strt := strings.Replace("ilovegolang", "o", "$$$", 1)
	println(strt)
	//字符串分割
	fmt.Println(strings.Split("i-love-go", "-"))
	fmt.Println(strings.ToLower("Go"))
	fmt.Println(strings.ToUpper("Go"))
	fmt.Println(strings.TrimSpace(" b  Go   "))
	fmt.Println(strings.Trim("**Go***", "*"))
	fmt.Println(strings.HasPrefix("Go", "G"))
	fmt.Println(strings.HasSuffix("Go", "o"))
	//println(err)
}
