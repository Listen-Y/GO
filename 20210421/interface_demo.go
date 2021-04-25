package main

import (
	"flag"
	"fmt"
	"time"
)

func main1() {
	test := make(map[string]bool)
	test["a"] = true
	fmt.Println(test)
}

type Read interface {
	Reader(data []int) (int, error)
}

type Write interface {
	Writer(data []int) (int, error)
}

type ReadImpl interface {
	Read
	Write(data []string) (int, error)
}

/**
接口类型，其实就是定义其的一些方法

表达一个类型属于某个接口只要这个类型实现这个接口中的方法即可,
并且父类接收子类的时候是子类给到的是地址

可以将任意一个值赋给空接口类型。
*/

type myRead struct {
	id int
}

func (my *myRead) Reader(data []int) (int, error) {
	fmt.Println(data)
	return 1, fmt.Errorf("a")
}
func test() Read {
	my := myRead{}
	return &my
}

func main2() {
	var key interface{}
	key = true
	key = 1
	key = "a"
	fmt.Println(key)

	var impl ReadImpl
	data := []int{1, 2, 3}
	_, _ = impl.Reader(data)

}

//延迟程序
var period = flag.Duration("period", 5*time.Second, "sleep")

func main() {
	flag.Parse()
	time.Sleep(*period)
	fmt.Println("a")
}
