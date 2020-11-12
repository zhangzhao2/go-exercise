package main

import "fmt"

func main() {
	var Numbers[5] int = [5]int{1, 2, 3, 4, 5} //c语言:int a[5] = {1,2,3,4,5}
	fmt.Println(Numbers[3])
	//2.部分赋值
	Numbers2 := [5]int{1,2}
	fmt.Println(Numbers2[4])
	fmt.Println(Numbers2[1])
	//3.指定元素初始化
	Numbers3 := [5]int{2:3, 3:4}
	fmt.Println(Numbers3[1])
	fmt.Println(Numbers3[3])
	//4.通过初始化确定数组长度
	Numbers4 := [...]int{7, 8, 9}
	fmt.Println(Numbers4[0])
	fmt.Println(Numbers4[2])

}
