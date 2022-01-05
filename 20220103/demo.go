package main

import (
	"fmt"
	"sync"
	"time"
)

func main1() {

	go printHello()

	for i := 0; i < 10; i++ {
		fmt.Println("主线程: ", "hello ", i)
		time.Sleep(time.Second)
	}
	fmt.Println("主线程: over")
}

func printHello() {
	// 首先让协程休眠一秒
	time.Sleep(time.Second)
	for i := 0; i < 10; i++ {
		fmt.Println("协程: ", "hello ", i)
		time.Sleep(time.Second)
	}
	fmt.Println("协程: over")
}

var (
	retMap2 = make(map[int]uint64, 200)
	// 互斥锁
	lock sync.Mutex
)

func main2() {
	for i := 1; i <= 20; i++ {
		go calculation(i)
	}

	time.Sleep(5 * time.Second) // 休眠5秒等待协程执行

	// 这里也需要加锁，防止出现并发读写问题
	lock.Lock()
	fmt.Println(retMap)
	lock.Unlock()
}

func calculation1(num int) {
	var ret uint64 = 1 //防止数据放不下
	for i := 2; i <= num; i++ {
		ret = ret * uint64(i)
	}

	lock.Lock()
	retMap2[num] = ret
	lock.Unlock()
}

var (
	retMap = sync.Map{}
)

func main3() {
	for i := 1; i <= 20; i++ {
		go calculation(i)
	}

	time.Sleep(2 * time.Second)
	retMap.Range(func(key, value interface{}) bool {
		fmt.Println("key: ", key, "value: ", value)
		return true
	})
}

func calculation(num int) {
	var ret uint64 = 1 //防止数据放不下
	for i := 2; i <= num; i++ {
		ret = ret * uint64(i)
	}

	retMap.Store(num, ret)
}

func main4() {
	// 初始化
	var intCh chan int
	intCh = make(chan int, 1) // 有缓存管道，可以缓存一个单位的数据

	//写入数据
	intCh <- 10

	// 读出数据
	ret := <-intCh
	fmt.Println(ret)
}

func main5() {
	intCh := make(chan int, 3)

	fmt.Printf("intCh的值: %v, intCh的类型: %T, intCh的地址: %p\n", intCh, intCh, &intCh)
	// intCh的值: 0xc00007e000, intCh的类型: chan int, intCh的地址: 0xc000006028
	close(intCh)
}

func main6() {
	intCh := make(chan int, 3)
	intCh <- 1
	intCh <- 2
	intCh <- 3
	close(intCh)
	//for ch := range intCh {
	//	fmt.Println(ch)
	//}
	for true {
		v, ok := <-intCh
		if !ok {
			break
		}
		fmt.Println(v)
	}
}

var dataMap = make(map[int]uint64, 20)

func main7() {
	intCh := make(chan int, 20)
	exitCh := make(chan bool, 1)
	go writeData(intCh)
	go readData(intCh, exitCh)

	// 下面这俩种写法都可以，也没有什么大区别
	/*	for true {
			v, ok := <- exitCh
			if !ok {
				break
			}
		fmt.Println(v)
		}*/

	for ch := range exitCh {
		if ch {
			break
		}
	}

	fmt.Println(dataMap)
}

func writeData(intCh chan int) {
	for i := 1; i <= 20; i++ {
		intCh <- i
	}
	// 如果管道不再需要放入数据就及时关闭
	close(intCh)
}

func readData(intCh chan int, exitCh chan bool) {
	for true {
		num, ok := <-intCh
		if !ok {
			break
		}
		// 计算
		var ret uint64 = 1 //防止数据放不下
		for i := 2; i <= num; i++ {
			ret = ret * uint64(i)
		}
		dataMap[num] = ret
	}
	// 计算完成后退出
	exitCh <- true
	close(exitCh)
}
