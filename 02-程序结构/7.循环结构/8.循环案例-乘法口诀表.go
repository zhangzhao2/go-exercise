package main

import "fmt"

//乘法口诀
func main() {
	//1.:考虑一行展示
	for j := 1; j<=9; j++{//行
		for i := 1; i<=j; i++{
			fmt.Printf("%d*%d=%d\t", j, i, j*i)  //-t --tab
		}
		fmt.Printf("\n")
	}
}
