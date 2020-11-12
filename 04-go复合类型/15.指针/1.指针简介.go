package main

import "fmt"

func main() {
	var num int = 1
	Test(num)
	fmt.Println(num)
}

func Test(num int)  {
	num = 30
}