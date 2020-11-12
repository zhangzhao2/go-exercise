package main

import "fmt"

func main()  {
	/*rand.Seed(time.Now().Unix())
	arr := make([]int, 10)

	for i := 0; i<len(arr); i++ {
		arr[i] = rand.Intn(100)
	}
	fmt.Println(arr)*/
	s := []int{5,9,0,2,7}

	insertSort(s)
	fmt.Println(s)
}
func insertSort(num []int)  {
	var j int = 0
	for i := 1; i < len(num); i++{
		tmp := num[i]
		fmt.Println("i=",i)
		for j = i; j > 0 && tmp < num[j-1]; j-- {
			fmt.Println("j=",j)
			num[j] = num[j-1]
			fmt.Println("num=",num)
		}

		num[j] = tmp
		fmt.Println("num=",num)
	}
}