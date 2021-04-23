package main

import "fmt"

//如果要在可变参数的函数中传递数组，那么你这个数组的大小一定是没有明显规定的
func main1() {
	nums := []int{1, 2, 3, 4}
	nums2 := []int{2, 3, 4, 5, 6, 7, 8, 9}
	var nums3 []int
	show(nums...)
	show(nums2...)
	show(nums3...)
	show()
}

func show(num ...int) {
	for _, v := range num {
		fmt.Println(v)
	}
}

func main() {
	fmt.Println(sum())           // "0"
	fmt.Println(sum(3))          // "3"
	fmt.Println(sum(1, 2, 3, 4)) // "10"
	values := []int{1, 2, 3, 4}
	fmt.Println(sum(values...)) // "10"
}

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}
