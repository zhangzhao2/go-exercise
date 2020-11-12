package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	//组织一个udp 地质结构，指定服务器的IP+Port
	srvAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8006")
	if err != nil {
		fmt.Println("ResolveUDPAddr err:", err)
		return
	}
	fmt.Println("udp 服务器地址结构创建完成！")

	//创建通信用的socket
	udpConn, err := net.ListenUDP("udp", srvAddr)
	if err != nil {
		fmt.Println("ListenUDP err:", err)
		return
	}
	defer udpConn.Close()
	fmt.Println("udp 服务器通信socket 创建完成")

	//读取客户端发送的数据
	buf := make([]byte, 4096)

	//与01-UDP服务器.go不同的是，这个里采用并发
	for {
		//返回三个值，分别为：读取的字节数，客户端的地址，error
		n, cltAddr, err := udpConn.ReadFromUDP(buf) //主go 程读取客户端发送数据
		if err != nil {
			fmt.Println("ReadFromUDP err:", err)
			return
		}
		//模拟处理器
		fmt.Printf("服务器读到%v 的数据：%s\n", cltAddr, string(buf[:n]))

		go func() { //每有一个客户连接上来，启动一个go程写数据
			//提取系统当前时间
			daytime := time.Now().String() + "\n"

			//回 写数据给客户端
			_, err = udpConn.WriteToUDP([]byte(daytime), cltAddr)
			if err != nil {
				fmt.Printf("WriteToUDP err:", err)
				return
			}
		}()
	}

}
