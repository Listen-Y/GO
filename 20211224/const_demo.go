package main

import "fmt"

const (
	a = iota
	b
	c
	d
)

func main1() {
	var num = 10
	fmt.Println("num的地址", &num)

	var ptr *int
	ptr = &num
	fmt.Println("ptr的地址", &ptr)
	fmt.Println("ptr的数据", ptr)
	fmt.Println("ptr的数据表示的真实意义", *ptr)
}

func main2() {
	//var ptr *float64
	//var num = 10
	//ptr = &num // error 指针也有类型
}

func main3() {
	// 例一
	nums := [4]int{1, 2, 3, 4}
	for i := 0; i < 4; i++ {
		fmt.Println(nums[i])
	}

	// 例二
	i := 0
	for {
		if nums[i] == 3 {
			break
		}
		i++
	}

	for index, value := range nums {
		fmt.Println("index: ", index, "value: ", value)
	}

	parts := map[string]string{}
	parts = make(map[string]string)
	parts["a"] = "aa"
	parts["b"] = "bb"
	parts["c"] = "cc"
	for k, v := range parts {
		fmt.Println("k:", k, "v:", v)
	}
}

func main4() {
	var score int
	_, _ = fmt.Scan(&score)
	switch {
	case score > 90:
		fmt.Println("A")
	case score > 80:
		fmt.Println("B")
	case score > 70:
		fmt.Println("C")
	default:
		fmt.Println("D")
	}
}

func main() {

	var num = 10
	var name = "name"
	findType(num)
	findType(name)
}

func findType(a interface{}) {
	switch a.(type) {
	case string:
		fmt.Println("string")
	case int64:
		fmt.Println("int64")
	case int:
		fmt.Println("int")
	case byte:
		fmt.Println("byte")
	case bool:
		fmt.Println("bool")
	default:

	}
}

func test(name string) {
	switch name {
	case "a", "x":
		fallthrough
	case "b":
		fmt.Println("b")
	default:

	}
}
