package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("客户端建立...")

	dial, err := net.Dial("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("客户端请求连接失败", err)
		return
	}

	//功能一:客户端可以发送单行数据,然后就退出, 在这里获得输入的数据
	reader := bufio.NewReader(os.Stdin) //os.stdin代表标准输入[终端]
	//从终端读取一行用户输入,并准备发送给服务器
	line, err := reader.ReadString('\n') // 以换行为输入的结尾
	if err != nil {
		fmt.Println("ReadString err=", err)
	}
	//再将line发送给服务器
	n, err := dial.Write([]byte(line))
	if err != nil {
		fmt.Println("conn write err=", err)
	}
	fmt.Printf("客户端发送了%d字节的数据，并退出", n)
}
