package main

import "fmt"

func main() {
	var i interface{}
	i = "abc"
	value, ok := i.(string)
	if ok{
		fmt.Println(value)
	}else{
		fmt.Println("类型断言错误")
	}
}
