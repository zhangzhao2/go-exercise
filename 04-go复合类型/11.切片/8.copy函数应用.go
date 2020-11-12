package main

import "fmt"

func main() {
	s1 := []int{1, 2}
	s2 := []int{3, 4, 5, 6, 7}
	s3 := []int{10, 11}
	copy(s2, s1)
	fmt.Println(s2)
	copy(s3, s1)
	fmt.Println(s3)
}
