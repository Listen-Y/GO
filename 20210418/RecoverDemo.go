package main

import "fmt"

//如果在deferred函数中调用了内置函数recover，
//并且定义该defer语句的函数发生了panic异常，
//recover会使程序从panic中恢复，
//并返回panic value。
//导致panic异常的函数剩余逻辑不会继续运行，但能正常执行defer所修饰的函数并且正常返回。
//在未发生panic时调用recover，recover会返回nil。
func main() {

	fmt.Println(makePanic())
}

func makePanic() (err error) {
	fmt.Println("进入make函数")
	fmt.Println("定义defer函数")
	defer func() {
		fmt.Println("进入defer修饰的函数")
		if p := recover(); p != nil {
			err = fmt.Errorf("修复panic, %v\n", p)
		}
		fmt.Println("退出defer修饰的函数")
	}()
	fmt.Println("制造panic")
	panic("我是panic")
	fmt.Println("退出make函数")
	err = fmt.Errorf("我是err")
	return err
}
