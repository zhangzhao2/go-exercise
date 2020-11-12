package main

import (
	"fmt"
)

func main() {
	//用continue实现，计算1-100之间的除了能被7整除之外所有整数的和
	var sum int
	for i:=1;i<=100;i++{
		if i % 7 == 0{
			continue
		}
		sum += i
	}
	fmt.Println(sum)
}
