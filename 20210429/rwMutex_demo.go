package main

import (
	"fmt"
	"sync"
	"time"
)

/*
这个读写锁和Java中的readwritelock差不多，都是读写锁，都是用在读多写少的场景，也就是说多个读可以并行进行，但是写操作只能单独进行
RWMutex只有当获得锁的大部分goroutine都是读操作，
而锁在竞争条件下，也就是说，goroutine们必须等待才能获取到锁的时候，
RWMutex才是最能带来好处的。RWMutex需要更复杂的内部记录，
所以会让它比一般的无竞争锁的mutex慢一些。
*/

var rwLock sync.RWMutex
var dolor = 10

func main() {
	for i := 0; i < 10; i++ {
		go read()
		go write()
		go read()
	}
	time.Sleep(3 * time.Second)
	fmt.Println("=")
	read()
}

func read() {
	rwLock.RLock()
	defer rwLock.RUnlock()
	fmt.Println(dolor)
}

func write() {
	rwLock.Lock()
	defer rwLock.Unlock()
	dolor += 1
}

/*
对于go并发的一些特点
在一个独立的goroutine中，每一个语句的执行顺序是可以被保证的，也就是说goroutine内顺序是连贯的，也就是说不会有什么指令重排序之类的
而其在go中的mutex的同步代码块中所有变量的内存也是可见的
*/
