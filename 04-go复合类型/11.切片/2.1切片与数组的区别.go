package main

import "fmt"

func main() {
	//切片和数组的区别
	//数组[]里面的长度时固定的一个常量， 数组不能修改长度， len和cap永远都是5
	a := [5]int{}
	fmt.Printf("len = %d, cap = %d\n", len(a), cap(a))

	//切片， []里面为空，或者为...，切片的长度或容量可以不固定
	s := []int{}
	fmt.Printf("1: len = %d, cap = %d\n", len(s), cap(s))

	//综上：两个看出区别，就是s := []int{},中括号里面有没数字，而c语言里面解决空间大小的是malloc申请空间来实现，实现起来比较简单

	s = append(s, 11) //给切片末尾追加一个成员
	fmt.Printf("append: len = %d, cap = %d\n", len(s), cap(s))
}
