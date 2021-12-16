package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main1() {

	var num = 100
	var p *int
	p = &num

	pp := &p

	*p = 200

	//0xc042004028 0xc0420080e0 200 0xc0420080e0 200
	//所以指针其实也是有自己地址的，其存储的是一个他所指向的地址，那个地址里存储这真正的数据
	fmt.Println(&p, p, *p, &num, num)
	fmt.Println(*pp)

}

func main2() {
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

func main3() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	str := "abcdef"
	for index, data := range str {
		fmt.Println(str[index], data)
	}
}

func main4() {

	str := "abcdef哈佛g"
	for index, data := range str {
		fmt.Println(index, str[index], data)
	}

	fmt.Println("================")

	//这种方式会有乱码，因为一个汉字占用三个字节
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c \n", str[i])
	}

	fmt.Println("========")
	str2 := []rune(str)

	for data := range str2 {
		fmt.Printf("%d%c\n", data, str2[data])
	}
}

//
func main6() {
	var count = 0

	for {
		rand.Seed(time.Now().UnixNano())
		var num = rand.Intn(100) + 1
		if num == 99 {
			break
		}
		count++
	}

	fmt.Println(count)
}

//goto语句的使用
func main7() {
	fmt.Println(1)
	fmt.Println(2)
	goto label1
	fmt.Println(3)
	fmt.Println(4)
	fmt.Println(5)
label1:
	fmt.Println(6)
}

func main() {

}

func operator(a float64, b float64, operator byte) (result float64) {

	switch operator {
	case '+':
		result = a + b
	case '-':
		result = a - b
	case '*':
		result = a * b
	case '/':
		result = a / b
	default:
		fmt.Println("ERROR")
	}

	return
}
