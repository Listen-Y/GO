package main

import (
	"fmt"
	"log"
)

func main1() {
	//对于map和切面在go函数传递中就不是浅拷贝
	ids := map[string]int{"a": 0}
	changeMap(ids)
	fmt.Println(ids["a"])
}

func changeMap(ids map[string]int) {
	ids["a"] = 999
}

func main2() {
	nums := [...]int{1, 2, 3}
	changeNums(&nums)
	for _, v := range nums {
		fmt.Println(v)
	}
}

func changeNums(nums *[3]int) {
	nums[0] = 99
}

func main3() {
	fmt.Println(fib(9))
}

func fib(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return fib(n-2) + fib(n-1)
}

func main5() {

	result, temp := fun1(1, 2, 3, 4, 5, 76)
	fmt.Println(result, temp)
}

func fun1(nums ...int) (int, int) {
	var result = 0
	var temp = 0
	for i, v := range nums {
		fmt.Println(i, " ", v)
		result += v
		temp += i
	}
	return result, temp
}

//上面的代码可以写成这个样子
func fun2(nums ...int) (result, temp int) {
	for i, v := range nums {
		fmt.Println(i, " ", v)
		result += v
		temp += i
	}
	return
}
func main6() {
	var erro error
	is(erro)
}

func is(err error) {
	if err != nil {
		fmt.Println("OK")
	} else {
		fmt.Println("Bad")
	}
	_ = fmt.Errorf("parsing %s as HTML: %v", "www.baidu.com", err)
}

//匿名函数 add返回一个匿名函数。返回的该匿名函数都会double一下结果
func add(a int) func() int {

	return func() int {
		a++
		return a + a
	}
}

func squares(X int) func() int {
	var x = X
	return func() int {
		x++
		return x * x
	}
}

func main7() {

	//用result接受这个匿名函数，每次调用的时候都是调用这个匿名函数，然后执行匿名函数中的逻辑
	result := add(1)
	fmt.Println(result())
	fmt.Println(result())
}

//函数类型
type test func(int, int) int

func one(a int, b int) int {
	return 1 + 1
}

func two(a int, b int) int {
	return 2 + 2
}

func main8() {

	fun1 := one(1, 2)
	fmt.Println(fun1)
	fmt.Printf("%T\n", 1)
	fun1 = two(2, 2)
	fmt.Println(fun1)
	fmt.Printf("%T\n", fun1)

	fmt.Println("-------------")
	var name = 19
	fmt.Printf("%T\n", name)

	var fun2 test
	fun2 = one
	result := fun2(1, 1)
	fmt.Printf("%T\n", fun2)
	fmt.Println(result)

}

func main() {
	fmt.Println("--------")
	getLog()
	err := getErr()
	fmt.Println(err)

}

func getLog() {
	log.Println("我是日志")
}

func getErr() error {

	return fmt.Errorf("%s\n", "www.baiducom")
}
