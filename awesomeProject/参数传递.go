package main

import "fmt"

func main() {
	arr := []int{1, 2, 55, 66, 99, 88, 77}
	arr2 := [3]string{"sss", "uuu", "lll"}

	chun(arr2)
	fmt.Println("员arr2：", arr2)

	chun1(arr)
	fmt.Println("原arr是：", arr)

}
func chun(arr [3]string) {
	arr[2] = "999"
	fmt.Println("更改后arr2:", arr)
}
func chun1(arr []int) {
	arr[3] = 888888
	fmt.Println("更改后arr:", arr)
}
