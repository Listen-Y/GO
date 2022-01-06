package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("服务器开始监听...")

	// net.Listen("tcp","o.0.0.0:8888")
	// 1. tcp表示使用网络协议是tcp
	// 2. 0.0.0.0.0:8888 表示在本地监听8888端口
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen error")
		return
	}

	defer listen.Close() // 必须关闭

	// 循坏等待客户端来连接
	for true {
		fmt.Println("等待客户端连接...")
		accept, err := listen.Accept()
		if err != nil {
			fmt.Println("accept error: ", err)
		} else {

			//获取到连接进行逻辑处理
			fmt.Println("accept success: ", accept)

			//创建一个新的切片,用户读取客户端发送来的数据
			buf := make([]byte, 1024)

			//1．等待客户端通过conn发送信息
			//2．如果客户端没有write[发送]，那么就会就阻塞在这里
			n, err := accept.Read(buf) //从连接中读取
			if err != nil {
				fmt.Println("服务器的Read err=", err)
			}
			//3．显示客户端发送的内容到服务器的终端
			fmt.Println("接受到的数据为: ", string(buf[:n]))
		}
	}
}
