package main

func body(a int) {
	a = a + 5
	println(a)
}
func main() {
	var a int = 99
	defer body(a)
	a++
	println(a)

}
