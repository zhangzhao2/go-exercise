package main

import "fmt"

func main() {
	var num int = 10
	update(&num)
	fmt.Println("num=", num)
}

func update(p *int)  {
	*p = 30
}