package main

import (
	te "awesomeProject/db"
	"fmt"
)

func main() {
	te.Dbconn()
	fmt.Println("test")
}
