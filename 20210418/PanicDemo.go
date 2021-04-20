package main

import "fmt"

//不是所有的panic异常都来自运行时，
//直接调用内置的panic函数也会引发panic异常；
//panic函数接受任何值作为参数。
//当某些不应该发生的场景发生时，我们就应该调用panic。
func mainA() {

	fmt.Println("正常")
	panic("异常")
	fmt.Println("no")
}

func mainB() {
	num := 1
	Reset(&num)
}

func Reset(x *int) {
	if x == nil {
		panic("x is nil") // unnecessary!
	}
	x = nil
}
