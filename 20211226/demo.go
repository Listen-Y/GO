package main

import (
	"fmt"
	"os"
	"strings"
	"time"
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

func main3() {
	fmt.Println("sum", sum(10, 20))
	file, _ := os.Open("D:/a.txt")
	defer file.Close()
}

func main() {

	// 获取当前时间
	now := time.Now()
	fmt.Printf("now=%v, now Type=%T\n", now, now)

	// 时间戳
	fmt.Println("毫秒时间戳", now.Unix())
	fmt.Println("纳秒时间戳", now.UnixNano())

	// 获取年月日时分秒
	fmt.Println("年=", now.Year())
	fmt.Println("月=", now.Month())
	fmt.Println("日=", now.Day())
	fmt.Println("时=", now.Hour())
	fmt.Println("分=", now.Minute())
	fmt.Println("秒=", now.Second())

	// 格式化日期
	fmt.Printf("当前年月日%02d-%02d-%02d %02d:%02d:%02d\n", now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())

	// 还有一种方式
	fmt.Println(now.Format("2006/01/02 15:04:05")) // 这个时间格式是固定的
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format("15:04:05"))
}
