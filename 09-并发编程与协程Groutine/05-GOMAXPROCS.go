package main

import (
	"fmt"
	"runtime"
)

func main() {
	n := runtime.GOMAXPROCS(0)//将CPU设置为 单核
	fmt.Println("n=", n)

	for {
		go fmt.Print(0) //子 go程

		fmt.Print(1) //主 go程
	}
}
