package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"time"
)

type PathError struct {
	path       string
	op         string
	createTime string
	message    string
}

// 自定义返回体
func (p *PathError) Error() string {
	return fmt.Sprintf("path=%s \nop=%s \ncreateTime=%s \nmessage=%s", p.path,
		p.op, p.createTime, p.message)
}

func Open(filename string) error {

	file, err := os.Open(filename)
	if err != nil {
		// 返回自定义error
		return &PathError{
			path:       filename,
			op:         "read",
			message:    err.Error(),
			createTime: fmt.Sprintf("%v", time.Now()),
		}
	}

	defer file.Close()
	return nil
}

func main1() {
	err := Open("/Users/5lmh/Desktop/go/src/test.txt")
	// 判断error类型
	switch v := err.(type) {
	case *PathError:
		fmt.Println("get path error,", v)
	default:
		fmt.Println("other error,", v)
	}

}

func main2() {
	go func(s string) {
		for i := 0; i < 10; i++ {
			i = 1
			time.Sleep(1000)
			fmt.Println(s)
		}
	}("world")
	// 主协程
	for i := 0; i < 10; i++ {
		i = 1
		// 切一下，再次分配任务
		runtime.Gosched()
		time.Sleep(1000)
		fmt.Println("hello")
	}
}

func rest() {
	defer fmt.Println("B.defer")
	// 结束协程
	//runtime.Goexit()

	func() {
		defer fmt.Println("D.defer")
	}()

	defer fmt.Println("C.defer")
	fmt.Println("B")
}

func main3() {
	go func() {
		defer fmt.Println("A.defer")
		go func() {
			defer fmt.Println("B.defer")
			// 结束协程
			//runtime.Goexit()
			defer fmt.Println("C.defer")
			fmt.Println("B")
		}()

		fmt.Println("A")
	}()

	for {
	}
}

func a() {
	for i := 1; i < 10; i++ {
		runtime.Gosched()
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 1; i < 10; i++ {
		runtime.Gosched()
		fmt.Println("B:", i)
	}
}

func main4() {
	runtime.GOMAXPROCS(1)
	go a()
	go b()
	time.Sleep(time.Second)
}

type Job struct {
	Id      int
	RandNum int
}

type Result struct {
	Job *Job
	Sum int
}

func main5() {
	// 需要2个管道
	// 1.job管道
	jobChan := make(chan *Job, 128)
	// 2. result管道
	resultChan := make(chan *Result, 128)
	// 3. 创建工作池
	createWorkerPool(64, jobChan, resultChan)
	// 4.开个打印的协程
	go func(resultChan <-chan *Result) {
		// 遍历结果管道打印
		for result := range resultChan {
			fmt.Printf("job id:%v randnum:%v result:%d\n", result.Job.Id,
				result.Job.RandNum, result.Sum)
		}
	}(resultChan)

	var id = 0
	// 循环创建job，输入到管道
	for {
		job := Job{
			Id:      id,
			RandNum: rand.Int(),
		}
		jobChan <- &job
		id++
	}
}

// 创建工作池
// 参数1：开几个协程
func createWorkerPool(goroutineNum int, jobChan chan *Job, resultChan chan *Result) {
	for i := 0; i < goroutineNum; i++ {
		// 根据开协程个数，去跑运行
		go func(jobChan <-chan *Job, resultChan chan<- *Result) {
			// 执行运算
			// 遍历job管道所有数据，进行相加
			for job := range jobChan {
				// 随机数接过来
				r_num := job.RandNum
				// 随机数每一位相加
				// 定义返回值
				var sum int
				for r_num != 0 {
					tmp := r_num % 10
					sum += tmp
					r_num /= 10
				}
				// 想要的结果是Result
				r := &Result{
					Job: job,
					Sum: sum,
				}
				//运算结果扔到管道
				resultChan <- r
			}
		}(jobChan, resultChan)
	}
}

func main6() {
	test4()
	for {

	}
}
func test1() {
	// 1.timer基本使用 延时俩秒
	timer1 := time.NewTimer(2 * time.Second)
	t1 := time.Now()
	fmt.Printf("t1:%v\n", t1)
	t2 := <-timer1.C
	fmt.Printf("t2:%v\n", t2)
}
func test2() {
	// 2.验证timer只能响应1次
	timer2 := time.NewTimer(time.Second)
	for {
		// 第一次考验正常执行，第二次就会发生deadlock
		<-timer2.C
		fmt.Println("时间到")
	}
}
func test3() {
	// 3.timer实现延时的功能
	//(1)
	time.Sleep(time.Second)
	//(2)
	timer3 := time.NewTimer(2 * time.Second)
	<-timer3.C
	fmt.Println("2秒到")
	//(3)
	<-time.After(2 * time.Second)
	fmt.Println("2秒到")
}
func test4() {
	// 4.停止定时器
	timer4 := time.NewTimer(2 * time.Second)
	go func() {
		fmt.Println("进入func")
		<-timer4.C
		fmt.Println("定时器开始执行")
	}()
	//time.Sleep(2*time.Second)
	b := timer4.Stop()
	if b {
		fmt.Println("timer4已经关闭")
	} else {
		fmt.Println("timer4无法关闭")
	}
}
func test5() {
	// 5.重置定时器
	timer5 := time.NewTimer(3 * time.Second)
	timer5.Reset(1 * time.Second)
	fmt.Println(time.Now())
	fmt.Println(<-timer5.C)
}

func test6() {

}

func main() {
	// 1.获取ticker对象
	ticker := time.NewTicker(1 * time.Second)
	i := 0
	// 子协程
	go func() {
		for {
			//<-ticker.C
			i++
			fmt.Println(<-ticker.C)
			if i == 5 {
				//停止
				ticker.Stop()
			}
		}
	}()

	for {
	}
}
