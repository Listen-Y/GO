package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main1() {

	var s = ""
	var seq = ""

	//os提供了一些与操作系统交互的函数与变量
	for i := 0; i < len(os.Args); i++ {
		s += seq + os.Args[i]
		seq = " "
	}
	fmt.Println(s)
	//strings下有很多的工具
	fmt.Println(strings.Join(os.Args[0:], " "))

	for index, data := range os.Args[0:] {
		fmt.Println(index, " ", data)
	}
}

func main2() {
	//fmt.Println(fib(3))
	//fun()
	comple := getComplex(3, 4)
	fmt.Println(real(comple))
	fmt.Println(imag(comple))
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

func fun() {
	var err error
	_, err = os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
	log.Fatal("os.Getwd")
}

func getComplex(real int, imag int) complex128 {
	return complex(float64(real), float64(imag))
}

//常量
//可以自己使用type定义一个类型，然后在下面使用，格式位type name 类型
type mouth int16

const DAY mouth = 0
const TIMEOUT = 5 * DAY

//及时中间有常量改了值，iota还是按行自增
const (
	one   = iota
	two   = 9
	three = iota
)

func main() {
	//输出位mouth类型
	fmt.Printf("%T\n", DAY)
	fmt.Printf("%T\n", TIMEOUT)
	fmt.Println(one, two, three)
}
