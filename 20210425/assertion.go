package main

import (
	"fmt"
	"go/types"
)

/*
断言
断言是判断一个接口类型是不是nil，或者判断一个接口所存储的值是不是属于某个类型

他有俩重写法，一种是用一个参数来接收结果，这种情况一旦判断为nil或者判断类型不匹配会触发panic异常
而另外一种用俩个参数来接收就不会触发异常，而其返回false
也就说被断言判断的是一定是一个接口

而用于判断断言的条件那个参数，可以是接口，也可以是非接口
如果是接口那么被断言的要实现这个接口才能断言成功，如果是非接口，那么想要断言合法，就得非接口类型实现接口类型了


1.类型断言格式为：x.(T);
2.类型断言的必要条件就是x是接口类型，非接口类型的x不能做类型断言；
3.T可以是非接口类型，如想断言合法，则T必须实现x的接口；
4.T也可以是接口，则x的动态类型也应该实现接口T；
*/
func main() {

	var test interface{} = 10
	fmt.Printf("%T\n", test)
	val, ok := test.(int)
	fmt.Printf("%T\n", test)
	fmt.Println(val, ok)

	var test1 interface{}
	val, ok = test1.(int)
	fmt.Println(val, ok)

}

var data interface{} = 10

func main2() {
	//断言在switch中的使用
	val := data
	switch val.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case types.Nil:
		fmt.Println("nil")
	default:
		fmt.Println("other")
	}
}

func main3() {

}
