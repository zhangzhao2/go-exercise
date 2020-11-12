package main

import "fmt"

func main() {
	nums := [10]int{1,2,3,4,5,6,7,8,9,10}
	var p *[10]int
	var i int
	p = &nums
	fmt.Println(*p)
	fmt.Println((*p)[3])//[]优先级大于*i
	fmt.Println(p[0])
	for i=0; i<len(p); i++ {
		fmt.Println(p[i])
	}
	updatea(p)
	fmt.Println(nums)
}

func updatea(p *[10]int)  {
	p[0] = 100
}