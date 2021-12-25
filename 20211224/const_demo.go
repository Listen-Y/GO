package main

import "fmt"

func init() {
	fmt.Println("other init")
}

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

func main5() {

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

func add(a, b int) {
	sum := a + b
	fmt.Println(sum)
}

func main6() {
	a1 := 10
	b1 := 20
	add(a1, b1)
}

func sum(nums ...int) {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	fmt.Println(sum)
}

func sum1(arr ...int) {
	sum := 0
	for i, data := range arr {
		fmt.Println("下标：", i)
		sum += data
	}
	fmt.Println(sum)
}

func main11() {
	sum(1, 2, 3)
	sum(4, 5, 6, 7, 89, 0)
}

func Test(nums ...int) {
	var s []int
	s = nums
	s = append(s, 1)
}

func aFunction() (int, int, float64, string) {
	return 1, 1, 1.0, "a"
}

func main7() {
	m, _, f, s := aFunction()
	fmt.Println(m, f, s)
	fmt.Println(Test1())
}

// 这种不会修改源数据->
func Test1() (a int, b float32) {
	defer func(a int, b float32) {
		a = 99
		b = 99.9
	}(a, b)
	a = 10
	b = 2.1
	return a + 1, b + 2
}

// 这种会修改源数据
func Test2() (a int, b float32) {
	defer func() {
		a = 99
		b = 99.9
	}()
	a = 10
	b = 2.1
	return a + 1, b + 2
}

func main8() {
	fmt.Println(Test1())
	fmt.Println(Test2())
}

func namedMinMax(x, y int) (min, max int) {
	if x > y {
		min = y
		max = x
	} else {
		min = x
		max = y
	}
	return
}

func minMax(x, y int) (min, max int) {
	if x > y {
		min = y
		max = x
	} else {
		min = x
		max = y
	}
	return min, max
}

func Swap(a, b *int) {
	*a, *b = *b, *a
	fmt.Println(*a, *b)
}

func main9() {
	a := 10
	b := 20
	fmt.Println(a, b)
	Swap(&a, &b)
	fmt.Println(a, b)
}

func swap(a *int, b *int) {
	fmt.Println(a)
	c := a // 此时c也是指针类型
	fmt.Printf("%T, %v\n", c, c)
	*a = *b
	*b = *c
}

func main10() {
	f := swap // 函数类型推导
	fmt.Printf("%T, %v\n", f, f)
	var aa = 10
	var b = 20
	fmt.Println(aa, b)
	f(&aa, &b)
	fmt.Println(aa, b)
}

// type可以定义类型, 也可以给已存在的类型起别名
type fType func(a, b *int)
type fType1 func(a, b int) int
type fType2 func(a, b, c int) (int, string)
type fType3 int

func test1(a, b *int) {
	fmt.Println("test1")
}

func test2(a, b int) int {
	fmt.Println("test2")
	return 0
}

func test3(a, b, c int) (int, string) {
	fmt.Println("test3")
	return 0, "test3"
}

func test4(a, b, c int) (int, string) {
	fmt.Println("test4")
	return 0, "test4"
}

func main() {
	a := 1
	b := 2
	c := 3

	var f fType
	f = test1
	f(&a, &b)

	f1 := test2
	_ = f1(a, b)

	var f2 fType2
	f2 = test3
	_, _ = f2(a, b, c)
	f2 = test4
	_, _ = f2(a, b, c)
}
