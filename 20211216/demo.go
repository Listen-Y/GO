package main

import "fmt"

func main() {
	var chanInt chan int
	chanInt = make(chan int, 2)
	chanInt <- 1
	chanInt <- 2
	//(chanInt)
	for i := 0; i < 4; i++ {
		num, ok := <-chanInt
		fmt.Println(num)
		// 如果在一个已经关闭的管道里读取数据，并且此时缓冲区没有数据，那么会返回默认值
		// 如果这个管道没有关闭，此时管道缓冲区也没有数据，编译器检查也没有协程往这个管道加数据，那么就会发生死锁deadklock
		fmt.Println(ok)
	}
}
