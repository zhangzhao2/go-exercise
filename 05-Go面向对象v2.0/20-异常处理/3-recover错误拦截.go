package main

import "fmt"

func main() {
	Test2(11)
}

func Test2(n int) {
	defer TestRecover()
	var num [10] int
	num[n] = 12
	fmt.Println(num)
	fmt.Println("aaa")
}
func TestRecover() {
	fmt.Println(recover())
}