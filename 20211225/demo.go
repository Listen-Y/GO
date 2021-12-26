package main

import "fmt"
import "./other"

func init() {
	fmt.Println("init")
}

func main1() {
	fmt.Println("main")
	other.Other()
}

func fun1(a int) int {
	fmt.Println("a + a")
	return a + a
}

func fun2(a, b int) int {
	fmt.Println("a * b")
	return a * b
}

func fun3(f1 func(a int) int, f2 func(a, b int) int, a, b int) int {
	return f1(a) + f2(a, b)
}

func main() {
	f1 := fun1
	f2 := fun2
	fmt.Println(fun3(f1, f2, 1, 2))

	// 第二种用法
	fmt.Println(fun3(func(a int) int {
		return a + a
	}, func(a, b int) int {
		return a * b
	}, 1, 2))
}
