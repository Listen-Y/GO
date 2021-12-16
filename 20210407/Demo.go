package main

import (
	a "./util"
	"fmt"
)

func main1() {

	var num1 = 10
	var num2 = 20

	fmt.Println(num1, num2)
	ret := a.Operator(num1, num2)
	fmt.Println(ret)
	a.Ha()

	result := a.FbNum(50)
	fmt.Println(result)
}

func addPro(fun func(int, int) int, a int, b int) int {
	return fun(a, b)
}

func main2() {
	var num = addPro(add, 10, 20)
	fmt.Println(num)
}

func add(a int, b int) int {
	return a + b
}

func add2(a int, b int) int {
	return a - b
}

type myAdd func(int, int) int

func main3() {
	var f myAdd
	f = add
	ret := f(10, 20)
	fmt.Println(ret)
	f = add2
	ret = f(10, 20)
	fmt.Println(ret)
}

func mord(a int, nums ...int) (sum int) {
	sum += a
	for index, data := range nums {
		sum += data
		fmt.Println(index)
	}
	return
}

func main4() {
	ret := mord(10, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(ret)
}

//a和b只是也是存储的一个地址，下面交换只是交换了a和b的地址
//而*a和*b才是在交换值
func swap(a *int, b *int) {
	fmt.Println(a)
	c := a
	fmt.Printf("%T, %v\n", c, c)
	a = b
	b = c
}

func main5() {
	f := swap
	var aa = 10
	var b = 20
	fmt.Println(aa, b)
	f(&aa, &b)
	swap(&aa, &b)
	fmt.Println(aa, b)
}

var v1 = test()

func test() int {
	fmt.Println("AAA")
	return 90
}

func init() {
	fmt.Println("BBB")
}

func main() {
	fmt.Println("CCC")
}
