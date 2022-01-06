package sharedRPC

// 共享结构体
type MyFloats struct {
	A1, A2 float64
}

// 共享接口
type MyInterface interface {
	Multiply(arguments *MyFloats, reply *float64) error
	Power(arguments *MyFloats, reply *float64) error
}

/*
sharedRPC包定义了一个名为MyInterface的接口和一个名为MyFloats 的结构，
客户端和服务器都将会使用到。
只有RPC服务器需要实现这个接口。
*/
