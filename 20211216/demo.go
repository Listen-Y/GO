package main

import (
	"fmt"
	"reflect"
)

func main1() {
	var chanInt chan int
	chanInt = make(chan int, 2)
	chanInt <- 1
	chanInt <- 2
	close(chanInt)
	for i := 0; i < 4; i++ {
		num, ok := <-chanInt
		fmt.Println(num)
		// 如果在一个已经关闭的管道里读取数据，并且此时缓冲区没有数据，那么会返回默认值
		// 如果这个管道没有关闭，此时管道缓冲区也没有数据，编译器检查也没有协程往这个管道加数据，那么就会发生死锁deadklock
		fmt.Println(ok)
	}
	num := 109
	typeOf := reflect.TypeOf(num)
	name := typeOf.Name()
	fmt.Println(name)
}

func main() {
	var num = 10
	//value := reflect.ValueOf(num)
	//value.SetInt(20) //这样操作会报错,如果要进行set,必须处理的是由指针带来的value
	valueOf := reflect.ValueOf(&num)
	valueOf.Elem().SetInt(20)

	fmt.Println(num)
}
