package main

import "fmt"

func main() {
	var num1 [5]int = [5]int{1, 2, 3, 4, 5}
	var num2 [5]int = [5]int{3, 2, 3, 4, 5}
	b := compareValue(num1, num2)
	if b {
		fmt.Println("数组一致")
	}else {
		fmt.Printf("数组不一致")
	}

	//fmt.Println()
}

func compareValue(n1 [5]int, n2 [5]int) bool {
	var b bool = true
	//1：判断两个数组的长度是否一致
	if len(n1) == len(n2){
		//2:判断值是否一致
		for i := 0; i < len(n1); i++ {
			if n1[i] == n2[i]{
				continue
			}else {
				fmt.Printf("i=%d\n", i)
				b = false
				break
			}
		}
	}else{
		b = false
	}
	return b
}
