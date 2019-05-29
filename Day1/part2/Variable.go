package main

import "fmt"

func main() {
	var a = "This is a"
	fmt.Println(a)

	b := "This is b"  // :은 뒤에 있는 변수유형에 따라 초기화됨
	b = "This is c"
	fmt.Println(b)	
}