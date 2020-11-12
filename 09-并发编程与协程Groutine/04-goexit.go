package main

import (
	"fmt"
	"runtime"
)

func test() {
	defer fmt.Println("ccccccccccc")
	runtime.Goexit()//退出当前go程序
	defer fmt.Println("ddddddd")
}

func main() {
	go func() {
		defer fmt.Println("aaaaaa")
		test()
		fmt.Println("bbbbbbb")
	}()
	for{
		;
	}
}
