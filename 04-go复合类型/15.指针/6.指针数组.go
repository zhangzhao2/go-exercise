package main

import "fmt"

func main()  {
	var p[2] *int
	var a int = 5
	var b int = 6
	p[0] = &a
	p[1] = &b
	fmt.Println(p)//[0xc00008a010 0xc00008a018]
	fmt.Println(len(p))
	for k, value := range p{
		fmt.Println("value=", *p[k])
		fmt.Println("*value=", *value)
	}
}
