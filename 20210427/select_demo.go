package main

import (
	"fmt"
	"time"
)

/*
select语法上和switch很像，但是select的case是对于一个个channel的
select默认会阻塞，找到那个case的channel可以走通，就会进行那个代码块中执行逻辑
如果有多个channel好了那就会随便找一个去执行
如果没有channel好就会走的default中的代码，如果没有default代码，就会一直阻塞着
我们可以利用其实现轮询，就是我们再default中啥也不干，但是在select代码外加上一层for循环来实现轮询select语句

*/
func main() {

	chan1 := make(chan string)
	chan2 := make(chan string)

	go putChan1(chan1)
	go putChan2(chan2)

	for i := 0; i < 100; i++ {
		select {
		case str := <-chan1:
			fmt.Println(str)
		case str := <-chan2:
			fmt.Println(str)
		default:
			fmt.Println("轮询遇到default")
			time.Sleep(1 * time.Second)
		}
	}

}

func putChan1(put chan<- string) {
	for i := 0; i < 10; i++ {
		time.Sleep(300 * time.Millisecond)
		put <- fmt.Sprintf("%s_%d", "putChan1", i)
	}
	fmt.Println("-------------------------close chan1")
	close(put)
}

func putChan2(put chan<- string) {
	for i := 0; i < 10; i++ {
		time.Sleep(20 * time.Millisecond)
		put <- fmt.Sprintf("%s_%d", "putChan2", i)
	}
	fmt.Println("-------------------------close chan2")
	close(put)
}
