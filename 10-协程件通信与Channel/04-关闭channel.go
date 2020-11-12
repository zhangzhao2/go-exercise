package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 0)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()
	time.Sleep(time.Second * 2)

	for num := range ch{
		fmt.Println("num=", num)
	}
}
