package main

func main() {
	var a, b int = 1, 2
	c, d := 3, 4
	var e, f = 123, 99

	e, _ = f, e
	_, b = 4, 7
	println(a, b, c, d, e, f)
	_, n, m := number()
	println(n, m)

}
func number() (int, int, string) {
	a, b, c := 1, 2, "qian"
	return a, b, c
}
