package main

import (
	"fmt"
	"os"
	"strings"
)

func addUpper() func(i int) int {
	num := 10
	return func(i int) int {
		num = num + i
		return num
	}
}

func main1() {
	f := addUpper() // 注意这里先进行调用addUpper函数, 然后将返回值进行函数类型推导给了f
	fmt.Println(f(10))
	fmt.Println(f(10))
	fmt.Println(f(10))
}

func makeSuffix(suffix string) func(src string) string {
	return func(src string) string {
		if !strings.HasSuffix(src, suffix) {
			return src + suffix
		}
		return src
	}
}

func main2() {
	f := makeSuffix(".jpg")

	fmt.Println(f("a.jpg"))
	fmt.Println(f("b"))
	fmt.Println(f("c"))
}

func sum(a, b int) int {
	defer func(a int) {
		a++
		fmt.Println("a=", a)
	}(a)
	defer func(b int) {
		b++
		fmt.Println("b=", b)
	}(b)
	ret := a + b
	a++
	b++
	fmt.Println("ret=", ret)
	return ret
}

func main() {
	fmt.Println("sum", sum(10, 20))
	file, _ := os.Open("D:/a.txt")
	defer file.Close()
}
