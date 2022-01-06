package main

import (
	"fmt"
	"log"
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

func main4() {
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

type Monster struct {
	Name  string `myJson:"monster_name"`
	Age   int    `myJson:"monster_age"`
	Score float32
	Sex   string
}

func (s Monster) Print() {
	fmt.Println(" ---start----")
	fmt.Println(s)
	fmt.Println("---end----")
}

func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

func testStruct(a interface{}) reflect.Value {
	valueOf := reflect.ValueOf(a)
	kind := valueOf.Kind()
	// 判断是否为结构体
	if kind != reflect.Struct {
		log.Fatal("not a struct")
	}
	return valueOf
}

func main() {
	monster := Monster{
		Name:  "jack",
		Age:   21,
		Score: 100,
		Sex:   "man",
	}
	valueOf := testStruct(monster)
	typeOf := reflect.TypeOf(monster)

	// 获得结构体字段数目
	numField := valueOf.NumField()
	fmt.Println("Monster has ", numField, "  fields")

	// 获得所有结构体字段,并遍历
	for i := 0; i < numField; i++ {

		fieldValue := valueOf.Field(i)
		fieldType := typeOf.Field(i)
		fmt.Println("Field ", i, " type is ", fieldType)
		fmt.Println("Field ", i, " value is ", fieldValue)

		// 获得tag数据
		get := typeOf.Field(i).Tag.Get("myJson")
		if get != "" {
			fmt.Println("Field ", i, " myJson is ", get)
		}

	}
	fmt.Println("====================================")

	// 获得方法数据, 如果方法是小写的,这里是统计不到的
	numMethod := valueOf.NumMethod()
	fmt.Println("Monster has ", numMethod, "  methods")

	// 遍历方法
	for i := 0; i < numMethod; i++ {
		method := typeOf.Method(i)
		fmt.Println("Method ", i, " type is ", method)
	}
	fmt.Println("====================================")

	// 执行第一个方法 Print, Print方法没有参数,所以传nil
	valueOf.Method(1).Call(nil)

	// 调用执行有参数的方法 GetSum
	params := make([]reflect.Value, 0)
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(20))
	retValues := valueOf.Method(0).Call(params) // 传入和传出的参数都是value切片
	fmt.Println(retValues[0].Int())
}
