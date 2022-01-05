package main

import (
	"fmt"
	"reflect"
)

func main1() {
	num := 10

	// 获得valueOf
	valueOf := reflect.ValueOf(num)
	fmt.Println(valueOf.Kind()) // int

	// 通过valueOf获得typeOf
	i := valueOf.Interface()
	typeOf := reflect.TypeOf(i)
	fmt.Println(typeOf.Kind()) // int

	// 原获取到num
	num2 := i.(int)
	fmt.Println(num2)

	// 直接获取typeOf
	of := reflect.TypeOf(num2)
	fmt.Println(of.Kind()) // int
}

func main3() {
	// 通过反射计算和修改数据
	num := 10
	valueOf := reflect.ValueOf(num)

	// 计算数据
	i := valueOf.Int()
	fmt.Println(i + 10) // 20

	// 修改原数据, 注意这里传递的是地址
	valueOf = reflect.ValueOf(&num)
	// elem用于获取指针指向的变量
	valueOf.Elem().SetInt(20)
	fmt.Println(num) // 20
}

// 第一个参数funcPtr以接口的形式传入函数指针，
// 函数参数args以可变参数的形式传入，
// bridge函数中可以用反射来动态执行funcPtr函数
func bridge(functionPtr interface{}, args ...interface{}) {

}

//专门用于做反射
func test(b interface{}) {

	//1.如何将Interface}转成reflect.value
	rVal := reflect.ValueOf(b)
	//2.如何将reflect.Value-> interface{}
	ival := rVal.Interface()
	//3.如何将interface{}转成原来的变量类型,使用类型断言
	v := ival.(int)
	fmt.Println(v)
}

func main() {
	var v float64 = 1.2

	valueOf := reflect.ValueOf(v)
	fmt.Println(valueOf.Kind())
	fmt.Println(valueOf.Type())

	i := valueOf.Interface()
	f := i.(float64)
	fmt.Println(f == v)

	of := reflect.ValueOf(&v)
	of.Elem().SetFloat(2.4)
	fmt.Println(v)
}
