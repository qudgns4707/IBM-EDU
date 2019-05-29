package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	now := time.Now().Format("2006-01-02-15-04-05") //숫자들이 정해져잇음. 2006이 2006년이아니라 2006이라는것이 년도를 나타냄
	p(now)
}