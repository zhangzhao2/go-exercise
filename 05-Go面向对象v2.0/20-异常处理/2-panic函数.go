package main

import "fmt"

func main() {
	Test1(11)
}

func Test1(n int) {
	var num [10] int
	num[n] = 12
	fmt.Println("hhh")
	//程序员自己不会调用该函数，但是如果程序员自己写的程序中出现了比较严重的异常，那么系统会调用该函数，终止整个程序的执行。
	//panic("abc")
	//fmt.Println("hello")
}