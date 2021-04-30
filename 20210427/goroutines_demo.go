package main

import (
	"fmt"
	"time"
)

/*
协程

当一个程序启动时，其主函数即在一个单独的goroutine中运行，我们叫它main goroutine。
新的goroutine会用go语句来创建。
一个普通的函数或方法调用前加上关键字go。
go语句会使其语句中的函数在一个新创建的goroutine中运行。
而go语句本身会迅速地完成。
主函数返回时，所有的goroutine都会被直接打断，程序退出。
除了从主函数退出或者直接终止程序还有goroutine之间的通信之外，没有其它的编程方法能够让一个goroutine来打断另一个的执行，
但是之后可以看到一种方式来实现这个目的，
通过goroutine之间的通信来让一个goroutine请求其它的goroutine，
并让被请求的goroutine自行结束执行。
*/
func main() {

	start := time.Now().Unix()
	spinner(3 * time.Second)
	fib(45)
	end := time.Now().Unix()
	fmt.Println(end - start)

	start = time.Now().Unix()
	go spinner(3 * time.Second) //使用go关键字就可以让这个方法开启独立的一个协程去执行
	fib(45)
	end = time.Now().Unix()
	fmt.Println(end - start)

}

func spinner(delay time.Duration) {
	for _, val := range `-\|/` {
		fmt.Printf("\r%c", val)
		time.Sleep(delay)
	}
}

func fib(num int) int {
	if num < 2 {
		return 2
	}
	return fib(num-1) + fib(num-2)
}
