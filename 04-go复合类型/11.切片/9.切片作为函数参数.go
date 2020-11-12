package main

import "fmt"

func main() {
	s := make([]int, 10)
	Init(s)
	fmt.Println("s=", s)
}

func Init(num []int)  {
	for i := 0; i < len(num); i++ {
		num[i]= i
	}
}
