package main

import (
	"fmt"
	"sync"
	"time"
)

/*
mutex是go中的锁，其实和操作系统中的mutex差不多，就是一个互斥锁，并且是不可重入的
如果你想重入一个已经加锁的方法，那么就会造成死锁阻塞，解决办法就是将这个方法进行拆分，让重入的时候去另外一个方法，
但是另外的那个方法需要保证只有在同步代码块中才能调用，这个就需要我们自己做好注释和标识
再就是mutex有luck和unluck方法，一般我们将unlock用defer修饰做到解锁的万无一失，
虽然消耗会略微增大但是这点无足挂齿
*/

var lock sync.Mutex
var money = 0

func main1() {

	for i := 0; i < 10; i++ {
		go addMoney(1)
	}
	time.Sleep(3 * time.Second)
	fmt.Println(money)
}

func addMoney(add int) {
	lock.Lock()
	for i := 0; i < 10000; i++ {
		money += add
	}
	defer lock.Unlock()
}

func main() {
	add()
	time.Sleep(time.Second)
	fmt.Println(get())
}

//这样会造成死锁，就是因为go的锁不是一个可重入的锁，解决办法就是将需要重入的代码块做一个脱离
//比如下面将原本需要调用的add方法替换成addDemo，
func add() {
	lock.Lock()
	defer lock.Unlock()
	for i := 0; i < 10; i++ {
		money += 1
		addDemo()
	}
}

func addDemo() {
	for i := 0; i < 10; i++ {
		money += 1
	}
}

func get() int {
	lock.Lock()
	defer lock.Unlock()
	return money
}
