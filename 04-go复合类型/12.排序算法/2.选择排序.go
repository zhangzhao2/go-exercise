package main

import "fmt"

func main() {
	s := []int{5, 9, 0, 2, 7}

	for j := 0; j < len(s)-1; j++ {
		min := s[j]
		minIdenx := j
		for i := j+1; i < len(s); i++ {
			if min > s[i]{
				min = s[i]
				minIdenx = i
			}
		}
		if minIdenx != j {
			s[j], s[minIdenx] = s[minIdenx], s[j]
		}
	}
	fmt.Println(s)
}
