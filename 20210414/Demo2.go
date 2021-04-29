package main

import (
	"fmt"
)

const (
	A = 4
	B = 5
	C = 3
	D = 1
)

func main4() {

	var nums = [...]string{"a1", "s2", "a3"}

	/*
		通过这种方式给数组建立的映射关系和常量本身的值是无关的，而其其数组原有下标访问也是支持的
	*/
	money := [...]string{A: "aa", B: "bb", C: "cc", D: "dd"}

	for _, data := range nums {
		fmt.Println(data)
	}
	fmt.Println("--------")
	for i := 0; i < len(money); i++ {
		fmt.Println(money[i])
	}
	fmt.Println("--------")
	fmt.Println(money[A])
	fmt.Println(money[B])
	fmt.Println(money[C])
	fmt.Println(money[D])
}

//slice
func main5() {

	day := [...]string{"one", "two", "three", "four", "five"}

	head := day[:2]
	last := day[len(day)-2:]

	show(head, last)

	day[0] = "fakeOne"
	day[len(day)-1] = "fakeFive"
	show(head, last)

	//head截取的是前面部分的，如果在后面直接append是会影响原数组的，其实slice只是浅拷贝原始数组
	//如果容量够的话直接append在len位置后面，不管原始数组是否有数据
	head = append(head, "six")
	showSource(day)
	show(head, last)
	last = append(last, "seven")
	show(head, last)
	showSource(day)
	//如果添加的数据len大于容量cap的时候，slice就会开辟新的数组，
	//此时再去修改这个slice的数据就不会影响源数据
	last[0] = "fakeFour"
	show(head, last)
	showSource(day)

}

func showSource(day [5]string) {
	fmt.Println("***************")
	fmt.Println(day)
	fmt.Println("***************")
}

func show(head []string, last []string) {
	fmt.Printf("heead: len =%d, cap = %d\n", len(head), cap(head))
	fmt.Printf("last: len =%d, cap = %d\n", len(last), cap(last))
	for _, v := range head {
		fmt.Println(v)
	}
	for _, v := range last {
		fmt.Println(v)
	}
	fmt.Println("-------")
}

func main() {
	sre := make([]string, 4, 6)
	fmt.Println(len(sre), " ", cap(sre))

	//不指定容量的情况下，默认容量就是len
	sre1 := make([]string, 4)
	fmt.Println(len(sre1), " ", cap(sre1))
}
