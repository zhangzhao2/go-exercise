package main

import (
	"fmt"
	"time"
)

func main()  {
	ch := make(chan int)
	quit := make(chan bool)
	go func() {  //子go程序获取数据
		for{
			select {
				case num := <-ch:
					fmt.Println("num：", num)
				case <-time.After(3 * time.Second):
					quit <- true
					goto lable
			}
		}
	lable:
	}()
	for i:=0; i<2; i++{
		ch <- i
		time.Sleep(time.Second * 2)
	}
	<-quit //主go程，阻塞等待，子go通知，退出
	fmt.Println("finish!!!!")
}
