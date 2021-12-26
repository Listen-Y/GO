package main

import "fmt"
import "./other"

func init() {
	fmt.Println("init")
}

func main0() {
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

func main1() {
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

func main3() {

	res := 0

	func(a, b int) {
		fmt.Println(a + b)
		res = a + b
	}(10, 20)

	fmt.Println(res)

	res = func(a, b int) int {
		return a + b
	}(10, 20)
	fmt.Println(res)

	f1 := func(a, b int) int {
		return a * b
	}

	fmt.Println(f1(1, 2))
	fmt.Println(f1(2, 3))

}

var (
	F1 = func(a, b int) int {
		return a / b
	}
)

func main4() {
	fmt.Println(F1(4, 2))
}
