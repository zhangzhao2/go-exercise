package main

import "fmt"

func main() {

	Numbers := []int{11,12,13,14,15}

	for i, v := range Numbers{
		fmt.Printf("[%d]=%d\n", i, v)
	}
}
