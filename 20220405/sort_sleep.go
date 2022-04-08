package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	arr := []int{33, 99, 66, 44, 22}
	ch := make(chan int)
	wg := sync.WaitGroup{}
	startTime := time.Now()
	for _, v := range arr {
		wg.Add(1)
		go func(ch chan<- int, v int) {
			time.Sleep(time.Duration(v*100) * time.Millisecond)
			ch <- v
			wg.Done()
		}(ch, v)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	for v := range ch {
		fmt.Println(v)
	}
	fmt.Printf("耗时：%v\n", time.Now().Sub(startTime))
}
