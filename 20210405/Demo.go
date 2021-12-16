package main

import (
	"fmt"
)

/**
循环语句
*/

//if
func main1() {

	var a, b, c int
	_, _ = fmt.Scan(&a, &b, &c)
	if a > b {
		if a > c {
			fmt.Println("a重")
		} else if a == c {
			fmt.Println("a, c一样重")
		} else {
			fmt.Println("c重")
		}
	} else if a == b {
		if a > c {
			fmt.Println("a, b一样重")
		} else {
			fmt.Println("c重")
		}
	} else {
		if a < c {
			fmt.Println("c重")
		} else if b == c {
			fmt.Println("b, c一样重")
		} else {
			fmt.Println("b重")
		}
	}

}

//switch
func main2() {

	var score int
	fmt.Scan(&score)

	switch score > 700 {
	case true:
		fmt.Println("好")
	case false:
		fmt.Println("坏")
	default:
		fmt.Println("无")
	}

	//成绩评分
	switch score / 10 {
	case 10:
		fmt.Println("我来啦")
		fallthrough //意思为执行下一个结构, 该语句后面不能再有代码
	case 9:
		fmt.Println("A")
	case 8:
		fmt.Println("B")
	default:
		fmt.Println("C")
	}
}

//for循环
func main3() {
	i := 0
	for {
		if i > 10 {
			break
		}
		fmt.Println(i)
		i++
	}
	fmt.Println(i)
}

func main4() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", i, j, i*j)
		}
		fmt.Println()
	}
}

//for实现while 使用单for加continue注意死循环
func main5() {

	i := 0
	for {

		if i == 5 {
			break
		}
		if i%2 != 0 {
			i++
			continue
		}
		fmt.Println(i)
		i++
	}
}
