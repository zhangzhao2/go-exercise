package main

import "fmt"

func main() {
	fmt.Println("请输入比较数据的个数")
	var count int
	fmt.Scan(&count)

	//2.输入比较的数据，存入切片中
	s := make([]int, count)
	AddData(s)

	max := CompareValue(s)
	fmt.Println("max=", max)

}
func AddData(num []int) {
	for i := 0; i< len(num); i++{
		fmt.Printf("请输入第 %d 个数据\n", i+1)
		fmt.Scan(&num[i])
	}
}

func CompareValue(num []int) int {
	var max int = num[0]
	for i := 0; i<len(num);i++{
		if num[i] > max{
			max = num[i]
		}
	}
	return max

}