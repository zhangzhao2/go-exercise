package main

import "fmt"

func main() {
	//var nums[5] int = [5]int{1, 2, 3, 4, 5}  //c语言：int num[5] = {1,2,3,4,5}
	//nums := []int{} -> Numbers := []int{11,12,13,14,15}
	s := make([]int, 3, 5)//长度不能大于容量
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

}

