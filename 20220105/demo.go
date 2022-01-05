package main

import (
	"fmt"
	"time"
)

func main1() {
	intChan := make(chan int, 50)
	primeChan := make(chan int, 2000)
	exitChan := make(chan bool, 4)

	//写数据
	go writeData(intChan)
	// 读数据
	for i := 0; i < 4; i++ {
		go readData(intChan, primeChan, exitChan)
	}

	// 控制退出
	for i := 0; i < 4; i++ {
		<-exitChan
	}
	close(primeChan)
	close(exitChan)

	for num := range primeChan {
		fmt.Println(num)
	}

}

func writeData(intChan chan int) {
	for i := 0; i < 2000; i++ {
		intChan <- i
	}

	close(intChan)
}

func readData(intChan chan int, primeChan chan int, exitChan chan bool) {
	for num := range intChan {
		flag := true
		//判断num是不是素数
		for i := 2; i < num; i++ {
			if num%i == 0 { //说明该num不是素数
				flag = false
				break
			}
		}
		if flag {
			primeChan <- num
		}
	}
	//当数据取完
	exitChan <- true
}

func main2() {
	intChan := make(chan int, 2)
	testIn(intChan)
	//intChan <- 30
	testOut(intChan)
}

func testOut(outChan <-chan int) {
	fmt.Println(<-outChan)
	fmt.Println(<-outChan)
}

func testIn(inChan chan<- int) {
	inChan <- 10
	inChan <- 22
}

func main() {
	one := make(chan int, 1)
	two := make(chan int, 1)
	three := make(chan int, 1)
	go testPut(one, 1)
	go testPut(two, 2)
	go testPut(three, 3)

	for {
		flag := false
		select {
		case v := <-one:
			fmt.Println("one: ", v)
		case v := <-two:
			fmt.Println("two: ", v)
		case v := <-three:
			fmt.Println("three: ", v)
		default:
			fmt.Println("over")
			close(one)
			close(two)
			close(three)
			flag = true
		}
		if flag {
			break
		}
		time.Sleep(time.Second)
	}
}

func testPut(ch chan<- int, times int) {
	for i := 1; i < times*3; i++ {
		num := i * times
		ch <- num
	}
	//close(ch) // 使用elect的时候这里不能close, 不然select的对应channel的case处一直会返回零值
}
