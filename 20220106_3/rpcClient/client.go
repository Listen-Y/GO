package main

import (
	"./sharedRPC"
	"fmt"
	"net/rpc"
)

func main() {

	//注意，尽管RPC服务器使用TCP，但使用rpc.Dial()函数代替net.Dial()连接RPC服务器。
	c, err := rpc.Dial("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println(err)
		return
	}

	args := sharedRPC.MyFloats{
		A1: 16,
		A2: -0.5,
	}

	//RPC客户端和RPC服务器之间通过call()函数交换函数名，参数和函数返回结果，
	//而RPC客户端对函数的具体实现一无所知。
	var reply float64
	err = c.Call("MyInterface.Multiply", args, &reply)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Reply (Multiply): %f\n", reply)

	err = c.Call("MyInterface.Power", args, &reply)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Reply (Power): %f\n", reply)
}
