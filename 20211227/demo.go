package main

import "fmt"

func main1() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("捕获异常")
			fmt.Println(err)
		}
	}()
	num1 := 10
	num2 := 0
	fmt.Println(num1 / num2)
}

func main2() {
	defer func() {

		func() {
			err := recover()
			if err != nil {
				fmt.Println("捕获异常")
				fmt.Println(err)
			}
		}()

	}()
	num1 := 10
	num2 := 0
	fmt.Println(num1 / num2)
}

func test1() error {
	return fmt.Errorf("error")
}
func test2() {
	err := test1()
	if err != nil {
		panic(err)
	}
}

func main4() {
	var arr1 = [3]int{1, 2, 3}
	arr2 := [4]int{1, 2, 3, 4}
	arr3 := [...]int{1, 2, 3, 4, 5}
	arr4 := [...]string{1: "b", 2: "c", 0: "a"} // 通过下标定位index
	fmt.Println(arr1, arr2, arr3, arr4)

	// 方式一
	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}

	// 方式二
	for index, value := range arr1 {
		fmt.Println(index, value)
	}

	test3(&arr1)
}

func test3(arr *[3]int) {
	fmt.Println(arr)
}

func main5() {

	arr := [5]int{0, 1, 2, 3, 4}

	slice1 := arr[0:4] // 这底层就是分割出来一个切片, 采用左闭右的方式进行隔离出数据
	fmt.Println("slice", slice1)
	fmt.Println("len", len(slice1))
	fmt.Println("cap", cap(slice1))
	slice1[0] = 99
	fmt.Println("arr", arr) // 输出 arr [99 1 2 3 4] 这里看到slice1和arr使用的同一个内存空间
	slice1 = append(slice1, 4)
	slice1 = append(slice1, 5)
	slice1 = append(slice1, 6)
	slice1 = append(slice1, 7)
	slice1[1] = 999
	fmt.Println("arr", arr)      // 还是输出 arr [99 1 2 3 4] 这里看到slice1和arr使用的同一个内存空间
	fmt.Println("slice", slice1) // 输出 slice [99 999 2 3 4 5 6 7]
}

func main6() {

	/*	// 第一种方式
		arr := [5]int{0, 1, 2, 3, 4}
		_ := arr[0:4] // 这底层就是分割出来一个切片, 采用左闭右的方式进行隔离出数据

		// 第二种方式
		_ := make([]int, 0, 4)*/

	a := []int{1, 2, 3, 4, 5}
	s := make([]int, 1, 4)
	fmt.Println(s)
	copy(s, a)
	fmt.Println(s)
}

func main7() {
	arr := [5]int{1, 2, 3, 4, 5}

	var slice []int
	slice = arr[:]
	var slice2 = slice
	slice2[0] = 10
	fmt.Println("arr", arr)
	fmt.Println("slice", slice)
	fmt.Println("slice2", slice2)
}

func main8() {
	slice1 := make([]int, 1, 2)
	slice2 := make([]int, 1, 2)
	slice3 := make([]int, 1, 2)

	test5(slice1)
	test6(slice2)
	test7(slice3)

	fmt.Println("slice1", slice1)
	fmt.Println("slice2", slice2)
	fmt.Println("slice3", slice3)
}

func test5(slice []int) {
	slice[0] = 10
}

func test6(slice []int) {
	slice = append(slice, 5)
	slice[0] = 10
}

func test7(slice []int) {
	slice = append(slice, 5)
	slice = append(slice, 5)
	slice = append(slice, 5)
	slice = append(slice, 5)
	slice = append(slice, 5)
	slice = append(slice, 5)
	slice[0] = 10
}

func main9() {
	str := "abcdefg"
	//slice := str[4:]
	//mt.Println(slice)
	strArr := []byte(str)
	strArr[0] = 'm'
	fmt.Println("str:", str, "strArr:", string(strArr))

	str1 := "我是Listen"
	sreArr1 := []rune(str1)
	sreArr1[0] = '你'
	fmt.Println("strArr1:", string(sreArr1))
}
