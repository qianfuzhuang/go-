package main

import "fmt"

// 创建公共构造方法
type Animal struct {
	Name     string
	Sex      string
	Age      int
	function string
}

func (public Animal) shout(a, b, c string) {
	public.Name = a
	public.Sex = b
	public.function = c
	fmt.Printf("%s是%s的会%s", public.Name, public.Sex, public.function)
}
func (ag *Animal) Setage(age int) {
	if age <= 0 || age > 50 {
		println("年龄超出限制")
	} else {
		ag.Age = age
	}

}
func (ag *Animal) Getage() int {
	return ag.Age
}
func (an *Animal) eat() {
	fmt.Println("他会吃")
}

// 对Animal方法进行复用
type cat struct {
	Animal
}
type dog struct {
	Animal
}

func (c *cat) scret() {
	fmt.Println("我会抓人")
}
func main() {
	cats := &cat{}
	cats.Name = "波波"
	cats.Sex = "母"
	cats.Setage(56)
	cats.eat()
	cats.function = "喵星人"
	cats.shout("宗宗", "女", "睡觉")
	cats.scret()
	fmt.Println(*cats)
	fmt.Println("-----------------------------------------------")
	fmt.Println(cats.Getage())
	fmt.Println("--------------------dog")
	//	用dog进行复用
	dogs := dog{}
	dogs.Name = "旺财"
	dogs.Sex = "公"
	dogs.Setage(6)
	//dogs.Age = 2

	fmt.Println(dogs)

}
