package main

import "fmt"

func main() {
	//用continue计算1~100之间所有计数之和，除了能被7整除之外的所有整数的和
	var sum int
	for i:=1; i<100; i++{
		if i%7 ==0{
			continue  //假如为7，则不进行sum = sum +7操作，进行i=8的循环
		}

		sum += i
	}
	fmt.Println(sum)
}