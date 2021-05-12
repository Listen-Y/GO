package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

/*
unsafe包是一个采用特殊方式实现的包。
是由编译器实现的。
它提供了一些访问语言内部特性的方法，特别是内存布局相关的细节。
将这些特性封装到一个独立的包中，是为在极少数情况下需要使用的时候，使用unsafe包是不安全的。
此外，有一些环境因为安全的因素可能限制这个包的使用。
*/
func main1() {

}

/*
unsafe.Sizeof函数返回操作数在内存中的字节大小
如果是指针引用类型或包含引用类型的大小在32位平台上是4个字节，在64位平台上是8个字节。
*/
func main4() {
	man := per{1, "listen"}
	fmt.Println(unsafe.Sizeof(man))
}

/*
unsafe.Alignof 函数返回对应参数的类型需要对齐的倍数
通常情况下布尔和数字类型需要对齐到它们本身的大小（最多8个字节），其它的类型对齐到机器字大小。
*/

func main() {
	man := per{1, "listen"}
	fmt.Println(unsafe.Alignof(man))
}

/*
unsafe.Offsetof 函数的参数必须是一个字段 x.f，然后返回 f 字段相对于 x 起始地址的偏移量
包括可能的空洞。
*/

type per struct {
	Id   int
	Name string
}

func main3() {
	man := per{1, "listen"}
	fmt.Println(unsafe.Offsetof(man.Id))
	fmt.Println(unsafe.Offsetof(man.Name))
}

/*
一个普通的*T类型指针可以被转化为unsafe.Pointer类型指针，
并且一个unsafe.Pointer类型指针也可以被转回普通的指针，
被转回普通的指针类型并不需要和原始的*T类型相同
和普通指针一样，unsafe.Pointer指针也是可以比较的，并且支持和nil常量比较判断是否为空指针。
*/

/*
reflect包的DeepEqual函数可以对两个值进行深度相等判断。
DeepEqual函数使用内建的==比较操作符对基础类型进行相等判断
它可以工作在任意的类型上，甚至对于一些不支持==操作运算符的类型也可以工作
DeepEqual函数很方便，而且可以支持任意的数据类型，但是它也有不足之处。
例如，它将一个nil值的map和非nil值但是空的map视作不相等，
同样nil值的slice 和非nil但是空的slice也视作不相等
DeepEqual函数类似的地方是它也是基于slice和map的每个元素进行递归比较，
不同之处是它将nil值的slice（map类似）和非nil值但是空的slice视作不相等的值
*/

func main2() {
	fmt.Println(reflect.DeepEqual([]int{1, 2, 3}, []int{1, 2, 3}))        // "true"
	fmt.Println(reflect.DeepEqual([]string{"foo"}, []string{"bar"}))      // "false"
	fmt.Println(reflect.DeepEqual([]string(nil), []string{}))             // "false"
	fmt.Println(reflect.DeepEqual(map[string]int(nil), map[string]int{})) // "false"
}

/*
高级语言使得程序员不用再关心真正运行程序的指令细节，同时也不再需要关注许多如内存布局之类的实现细节。
因为高级语言这个绝缘的抽象层，我们可以编写安全健壮的，并且可以运行在不同操作系统上的具有高度可移植性的程序。

但是unsafe包，它让程序员可以透过这个绝缘的抽象层直接使用一些必要的功能，虽然可能是为了获得更好的性能。
但是代价就是牺牲了可移植性和程序安全，因此使用unsafe包是一个危险的行为。
*/
