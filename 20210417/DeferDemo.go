package main

import "fmt"

/*
resp.Body.close调用了多次，
这是为了确保title在所有执行路径下（即使函数运行失败）都关闭了网络连接。
随着函数变得复杂，需要处理的错误也变多，维护清理逻辑变得越来越困难。
而Go语言独有的defer机制可以让事情变得简单。

你只需要在调用普通函数或方法前加上关键字defer，
就完成了defer所需要的语法。
当执行到该条语句时，函数和参数表达式得到计算，但直到包含该defer语句的函数执行完毕时，
defer后的函数才会被执行，不论包含defer语句的函数是通过return正常结束，还是由于panic导致的异常结束。
你可以在一个函数中执行多条defer语句，它们的执行顺序与声明顺序相反。
也就是你后defer的语句先执行
*/

/**
总结下来就是：
1. defer语句表示延迟执行，也就是相当如Java的中的finally， 但是这个defer更灵活，不用try，只要所在函数执行完毕就会执行defer中的函数
2. 函数中后defer的语句先执行
3. defer还可以匿名函数
*/

//defer语句经常被用于处理成对的操作，
//如打开、关闭、连接、断开连接、加锁、释放锁。通过defer机制，
//不论函数逻辑多复杂，都能保证在任何执行路径下，资源被释放。
//释放资源的defer应该直接跟在请求资源的语句后。
func main() {

	fmt.Println(getDouble(2))
	fmt.Println(getDouble(3))

}

func getDouble(num int) int {
	defer func() { fmt.Printf("返回的结果是：%d\n", num*num) }()

	return num * num
}
