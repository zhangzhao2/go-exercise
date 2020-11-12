package main

import (
	"fmt"
)

func sing() {
	for i := 0; i < 50; i++ {
		fmt.Println("---正在唱, 隔壁泰山---")
	}
}

func dance() {
	for i := 0; i<50;i++{
		fmt.Println("---正在跳舞，赵四街舞---")
	}
}

func main() {
	go sing()
	go dance()
	for{
		;
	}

}