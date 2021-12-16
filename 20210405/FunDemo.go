package main

import (
	"fmt"
	"strconv"
)

func swap(a int, b int) {
	a, b = b, a
	fmt.Println(a, b)
}

func main10() {

	a := 10
	b := 20
	swap(a, b)
	fmt.Println(a, b)

}

func sum(nums ...int) {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	fmt.Println(sum)
}

//不定参数接受与遍历
func sum1(arr ...int) {
	sum := 0
	for i, data := range arr {
		fmt.Println("下标：", i)
		sum += data
	}
	fmt.Println(sum)
}

func main11() {
	sum1(1, 2, 3)
	sum1(4, 5, 6, 7, 89, 0)
}

//不定参数函数传递方式
func fun1(nums ...int) {
	fun2(nums[:]...)
	//左闭右开，0可以省略
	fun2(nums[0:2]...)
}
func fun2(nums ...int) {
	for data := range nums {
		fmt.Println(data)
	}
}

func fun3(a int, b int) int {
	result := a - b
	return result
}

func fun4(a int, b int) (ret int) {
	ret = a - b
	return
}

func main12() {
	fun3(10, 20)
	fun4(10, 20)
}

func fun5() (a int, b int, c int) {
	a = 10
	b = 20
	c = 30
	return
}

func main13() {
	num1, num2, _ := fun5()
	fmt.Println(num1, num2)
}

/**
函数类型
*/

func test1(a int, b int) (int, int) {
	fmt.Println(1)
	return a + b, a - b
}

func test2() {
	fmt.Println(2)
}

func test4() {
	fmt.Println(4)
}

func test3(a int, b int) (int, int) {
	fmt.Println(3)
	return a - b, a + b
}

type FUNTEST func(int, int) (int, int)
type FUNTEST1 func()

func main14() {
	var f FUNTEST
	f = test3
	ret, ret2 := f(10, 20)
	fmt.Println(ret, ret2)
	f = test1
	ret, ret2 = f(10, 20)
	fmt.Println(ret)
	fmt.Printf("%T", f)
}

var i = 0

func test6() {
	fmt.Println(i)
	i = 10
}

func test7() {
	fmt.Println(i)
}

func main17() {
	test6()
	test7()
	test5()
}

//变量的使用范围，如果程序中出现相同的变量名，优先使用自己内部的
//如果自己内部没有才会向外面去找，使用就近原则
func test5(num ...int) {
	i := 0
	for i := 0; i < len(num); i++ {
		fmt.Println(i)
	}
	fmt.Println(i)
	fmt.Println(i)
}

//字符串与基本数据类型的转换， 如果不支持转，就直接赋予默认值
func main18() {
	var num1 float64
	num1 = 129.42352352523
	var num2 = int8(num1)
	fmt.Println(num2)

	str := fmt.Sprintf("%f%d", num1, num2)
	fmt.Println(str)

	var is = true
	var num3 = 423.423532
	s := fmt.Sprintf("%t", is)
	fmt.Printf("%T, %v", s, s)
	ss := fmt.Sprintf("%s", str)
	fmt.Printf("%T, %v", ss, ss)
	sss := fmt.Sprintf("%f", num3)
	fmt.Printf("%T, %v", sss, sss)
}

func main() {
	is, _ := strconv.ParseBool("true")

	fmt.Printf("%T\n", is)

	var str = "1234"

	num, _ := strconv.ParseInt(str, 10, 8)
	fmt.Printf("%T, %v\n", num, num)

	var num1 int
	num1 = 10
	var num2 int32
	num2 = 10
	//在go中int的大小是和操作系统有关的，如果是32位就是4字节，如果是64位操作系统就是8字节
	fmt.Printf("%T\n", num1)
	fmt.Printf("%T", num2)

}
