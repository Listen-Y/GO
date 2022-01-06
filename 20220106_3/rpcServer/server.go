package main

import (
	"./sharedRPC"
	"fmt"
	"math"
	"net"
	"net/rpc"
)

// 服务端需要实现MyInterface 接口
type MyInterface struct{}

func (t *MyInterface) Multiply(arguments *sharedRPC.MyFloats,
	reply *float64) error {

	*reply = arguments.A1 * arguments.A2
	return nil
}
func (t *MyInterface) Power(arguments *sharedRPC.MyFloats,
	reply *float64) error {

	*reply = Power(arguments.A1, arguments.A2)
	return nil
}

func Power(x, y float64) float64 {
	return math.Pow(x, y)
}

func main() {

	myInterface := new(MyInterface)

	/*
		rpc.Register()函数的调用使这个程序成为RPC服务器。
		但是，由于RPC服务器使用TCP协议，
		它仍需要调用net.ResolveTCPAddr()和net.ListenTCP()。
	*/
	err := rpc.Register(myInterface)
	if err != nil {
		fmt.Println("register error", err)
		return
	}

	t, err := net.ResolveTCPAddr("tcp4", "0.0.0.0:8888")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := net.ListenTCP("tcp4", t)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}
		/*
			RemoteAddr()函数返回接入的RPC客户端IP地址和端口。
			rpc.ServerConn()函数为RPC客户端提供服务。
		*/
		fmt.Printf("%s\n", c.RemoteAddr())
		rpc.ServeConn(c)
	}
}
