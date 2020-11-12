package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("ziGO程序，i=", i)
			ch <- i
		}
	}()
	for i := 0; i < 5; i++ {
		num := <-ch
		fmt.Println("主go程序：", num)
	}
}
