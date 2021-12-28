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

func main() {

}
