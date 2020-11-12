package main

import (
	"fmt"
	"time"
)


func main()  {
	go func () {				// 创建一个 子go 程
		for i:=0; i<5;i++ {
			fmt.Println("------I'm goroutine -------")
			time.Sleep(time.Second)
		}
	}()

	fmt.Println("------I'm main-------")
}