package main

import (
	"bufio"
	"fmt"
	"image"
	"log"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var x int64
var lock sync.Mutex

func add1() {
	for i := 0; i < 5000; i++ {
		lock.Lock() // 加锁
		x = x + 1
		lock.Unlock() // 解锁
	}
}
func main1() {
	go add()
	go add()
	time.Sleep(5 * time.Second)
	fmt.Println(x)

	for {

	}
}

var wgroup sync.WaitGroup

func hello() {
	defer wgroup.Done()
	fmt.Println("Hello Goroutine!")
	time.Sleep(time.Second) // 假设这个任务耗时1秒
}
func main3() {
	wgroup.Add(1)
	go hello() // 启动另外一个goroutine去执行hello函数
	fmt.Println("main goroutine done!")
	wgroup.Wait()
	fmt.Println("goroutine done!")
}

var (
	shareData int
	wg        sync.WaitGroup
	rwLock    sync.RWMutex
)

func read() {
	defer wg.Done()
	rwLock.RLock()
	time.Sleep(10 * time.Millisecond)
	rwLock.RUnlock()
}

func write() {
	defer wg.Done()
	rwLock.Lock()
	shareData += 1
	time.Sleep(20 * time.Millisecond)
	rwLock.Unlock()
}

func main2() {
	start := time.Now()
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go write()
	}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(shareData)
	fmt.Println(end.Sub(start))
}

var icons map[string]image.Image

func loadIcons() {
	icons = map[string]image.Image{
		"left":  loadIcon("./left.png"),
		"up":    loadIcon("./up.png"),
		"right": loadIcon("./right.png"),
		"down":  loadIcon("./down.png"),
	}
}

func loadIcon(photoPath string) image.Image {
	// 打开文件
	file, err := os.Open(photoPath)
	if err != nil {
		log.Fatal("open file error")
		return nil
	}
	// 关闭
	defer file.Close()
	// 创建一个带缓冲区的Reader
	reader := bufio.NewReader(file)
	decode, _, _ := image.Decode(reader)
	return decode
}

var loadIconsOnce sync.Once

// Icon 是并发安全的
func Icon(name string) image.Image {
	loadIconsOnce.Do(loadIcons)
	return icons[name]
}

var m = make(map[string]int)

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}

func main4() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			set(key, n)
			fmt.Printf("k=:%v,v:=%v\n", key, get(key))
			wg.Done()
		}(i)
	}
	wg.Wait()
}

var syncMap = sync.Map{}

func main5() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			syncMap.Store(key, n)
			value, _ := syncMap.Load(key)
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()

	syncMap.Range(func(key, value interface{}) bool {
		fmt.Printf("key:%v value:%v\n", key, value)
		return true
	})
}

var num int64
var l sync.Mutex
var wGroup sync.WaitGroup

// 普通版加函数
func add() {
	// x = x + 1
	num++ // 等价于上面的操作
	wGroup.Done()
}

// 互斥锁版加函数
func mutexAdd() {
	l.Lock()
	num++
	l.Unlock()
	wGroup.Done()
}

// 原子操作版加函数
func atomicAdd() {
	atomic.AddInt64(&num, 1)
	wGroup.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 10000; i++ {
		wGroup.Add(1)
		go add() // 普通版add函数 不是并发安全的
	}
	wGroup.Wait()
	end := time.Now()
	fmt.Println(num)
	fmt.Println("1: " + end.Sub(start).String())

	num = 0
	start = time.Now()
	for i := 0; i < 10000; i++ {
		wGroup.Add(1)
		go mutexAdd() // 加锁版add函数 是并发安全的，但是加锁性能开销大
	}
	wGroup.Wait()
	end = time.Now()
	fmt.Println(num)
	fmt.Println("2: " + end.Sub(start).String())

	num = 0
	start = time.Now()
	for i := 0; i < 10000; i++ {
		wGroup.Add(1)
		go atomicAdd() // 原子操作版add函数 是并发安全，性能优于加锁版
	}
	wGroup.Wait()
	end = time.Now()
	fmt.Println(num)
	fmt.Println("3: " + end.Sub(start).String())
}
