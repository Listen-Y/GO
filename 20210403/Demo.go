//导入主函数的包
package main

import "fmt"

/**
基本语法
*/
func main1() {
	fmt.Println("aa")
}

func main2() {
	var a int = 10
	fmt.Print(a)

}

func main3() {
	PI := 3.1416926
	r := 2.5

	//自动推到类型，一般自动推到浮点数都是float64的
	//在go语言中，不同数据类型不能进行计算，可以通过类型转换解决
	s := r * r * PI
	l := 2 * PI * r

	fmt.Println("面积", s)
	fmt.Println("周长", l)
}

func main4() {
	a, b, c, d := 10, 20, 30, 40
	fmt.Println(a, b, c, d)
	e, f, g, h := 1, 2, 3, 4
	fmt.Println(e, f, g, h)
}

func main5() {
	//go语言的交换俩个数据
	// _ 表示匿名变量，在函数中会遇到，只接受一个变量，但是不能使用
	a, b := 10, 20
	a, b = b, a
	fmt.Println(a, b)
}

func main6() {
	b := 10
	c := 20.123423244234234
	fmt.Printf("b = %d\n", b)
	//默认保留6位小数
	fmt.Printf("c = %.3f\n", c)
}

func main7() {
	var is bool
	fmt.Println(is)
	is = true
	//%t表示输出一个布尔值
	fmt.Printf("%t\n", is)
	var str string
	str = "hahaha"
	fmt.Println(str)
}

func main8() {
	var str1, str2 string
	_, _ = fmt.Scan(&str1, &str2)
	//%p表示输出一个地址
	fmt.Printf("%p\n", &str1)
	fmt.Println(str1, str2)

}

func main9() {
	var str string
	//scan表示简单的数据，scanf表示格式化输入
	var num int
	fmt.Scanf("%s%d", &str, &num)
	fmt.Println(str, " ", num)
	/**
	变量命名
	1. 只能有数字和_和字母
	2. 开头只能是字母和_
	3. 区分大小写
	4. 不能使用关键词

	*/
}

func main10() {
	var num1 float32
	num1 = 1.123456789 //末尾会加一
	var num2 float64
	num2 = 1.23456789101112131415 //末尾会直接砍掉
	fmt.Println(num1, num2)
}

func main11() {
	var sre = "我是大王"
	//go语言中，一个汉字占三个字符，只要是为了和linux统一
	fmt.Println(len(sre))
}

func main12() {
	const (
		a    = iota
		b, e = iota, iota
		c    = iota
		d    = iota
		f, g = iota, iota
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
}

func main13() {
	num1 := 10.0
	var num2 = 20.0
	var ret = num1 / num2
	fmt.Println(ret)
}

func main14() {
	num1 := 10
	num2 := 20
	var ret = num1 % num2
	fmt.Println(ret)
}

func main15() {
	var num1 int8 = 12
	var num2 int = 13
	ret := int(num1) + num2
	fmt.Println(ret)
}

func main16() {
	var num1 = 10
	var num2 = 20
	num1 += num1
	num2 += num2
	b := num1 == num2
	fmt.Println(b)
}

func main17() {
	num := 10
	p := &num
	fmt.Println(*p)
	pp := &p
	fmt.Println(**pp)
}

func main18() {
	var a = 20
	var b = 30
	var c = 40
	c = c + b
	fmt.Println(a+b >= c && (c) >= a)
}

func main() {
	var year = 0
	fmt.Scan(&year)
	var is = year%400 == 0 || (year%4 == 0 && year%100 != 0)
	fmt.Println(is)
}
