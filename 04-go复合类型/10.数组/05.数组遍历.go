package main

import "fmt"

func main() {
	//从一个数组中取出最大的整数，最小整数，并且求总和，求平均值
	//1:定义数组
	var nums[5] int = [5]int{1,2,3,4,5}
	//2:定义两个变量存储最大值和最小值
	var max int = nums[0]
	var min int = nums[0]
	var sum int
	//3:循环数组和最大值和最下值比价
	for i := 0; i < len(nums); i++ {
		//每次循环将两个值比较将最大的值，存在max中。
		if nums[i] > max {
			max = nums[i]
		}
		if nums[i] < min {
			min = nums[i]
		}
		sum = sum + nums[i]
	}
	fmt.Println("最大值：",max)
	fmt.Println("最小值为", min)
	fmt.Println("和为：", sum)
	//这里为什么要变为float64, 因为除数，要精确小数点，所以用双精度。
	fmt.Printf("平均值为:%.2f", float64(sum)/5)

}
