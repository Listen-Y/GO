package main

import (
	"fmt"
	"time"
)

type Person struct {
	Id      int
	Name    string
	Address string
	Dob     time.Time
}

type man Person

func main5() {

	peo := man{10, "listen", "shanxi", time.Now().Local()}

	var peo1 man
	peo1.Address = "shanghai"
	fmt.Println(peo)
	//fmt.Println(peo1)
	//传递指针
	changPer(&peo)
	fmt.Println(peo)
	//对象中的某一个属性传递指针
	changeItem(&peo.Address)
	fmt.Println(peo)

}

func changeItem(str *string) {
	fmt.Println(str)
	*str = "aaa"
}

/*
对比上下俩个方法可以发现，如果传递的是基本数据类型的地址，那么指针存储的是地址，如果是结构体，指针存储的是 &+原对象
所以操作基本数据类型的时候得加*, 而操作结构体指针的时候不需要加*
*/

//在go中对于传给的结构体来说，如果你只是传给一个引用，那么需改还是不会改变源对象的，如果你传递的是一个地址，那么久不一样了
//Go函数的参数传递方式是值传递，这句话对结构体也是适用的
func changPer(per *man) {
	fmt.Println(per, " ", *per, " ", &per)
	per.Name = "999"
}

func main4() {
	nums := [...]int{1, 2, 3, 4}
	fmt.Println(nums)
	change(nums)
	fmt.Println(nums)
}

/*
在go中数组也算作基本数据类型，是放在栈区的
*/
func change(nums [4]int) {
	nums[0] = -1
	fmt.Println(nums)
}

func main() {
	//对于这种情况，指定下标的行为，只能是int整形，可以是乱序设置，但是其内部还是有序的，数组长度最大为设置的最大整数
	//对应下标没有设置的赋予初始值
	nums := [...]string{9: "aa", 2: "bb", 1: "cc"}
	fmt.Println(nums[3], " ", nums[0])

	for index, data := range nums {
		fmt.Println(index, " ", data)
	}
	fmt.Println(len(nums))
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
}
