package main

import (
	"fmt"
	"time"
)

/*
用两个channels将三个goroutine串联起来
第一个goroutine是一个计数器，用于生成0、1、2、……形式的整数序列，
然后通过channel将该整数序列发送给第二个goroutine；第二个goroutine是一个求平方的程序，
对收到的每个整数求平方，
然后将平方后的结果通过第二个channel发送给第三个goroutine
第三个再输出这三个数据的最大和最小值

channel主要分为用缓存和无缓存，还有双向和单向之分
无缓存的就是一个同步器，你放进去数据再放就会阻塞，你拿数据之后你再拿也会阻塞
有缓存也一样，就是可以多放几个数据，满了还会阻塞
双向就是你可以写也可以读，但是为了在方法中传参数的时候传递channel的时候就会区分出你只能输入或者只能输出，你在这个方法内部就是一个单向的
当然参数传递channel的时候是引用传递
再就是调用close方法关闭channel，如果关闭之后，在写入数据就会触发panic异常，但是这个channel还是可以读的，直到读到零值就会表示数据读取完了
所以我们一般只在写入完毕的时候关闭channel
channel可以是nil值，可以进行nil判断。还可以进行==比较，两个相同类型的channel可以使用==运算符比较。
如果两个channel引用的是相同的对象，那么比较的结果为真。一个channel也可以和nil进行比较。
*/
var chan1 = make(chan int)
var chan2 = make(chan int)

func main1() {

	go makeNums()
	go numNum()
	getNum()

}

func makeNums() {
	for i := 0; i < 3; i++ {
		chan1 <- i
	}
	close(chan1)
}

func numNum() {
	for val := range chan1 {
		time.Sleep(time.Second)
		chan2 <- val * val
	}
	close(chan2)
}

func getNum() {
	for val := range chan2 {
		fmt.Println(val)
	}
}

/*
改进版本
*/
func counter(put chan<- int) {
	for x := 0; x < 100; x++ {
		put <- x
	}
	close(put)
}

func squarer(put chan<- int, get <-chan int) {
	for v := range get {
		put <- v * v
	}
	close(put)
}

func printer(get <-chan int) {
	for v := range get {
		fmt.Println(v)
	}
}

func main2() {
	naturals := make(chan int)
	squares := make(chan int)
	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}

func main() {
	s := "a"
	//对于匿名函数，再最后的那个()里可以传入参数，再fun()的括号里接收参数，然后再匿名函数中使用这个参数
	go func(s string) {
		fmt.Println(s)
	}(s)
}
