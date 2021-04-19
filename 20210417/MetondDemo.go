package main

import "fmt"

func mainB() {
	p := Peo{1, "listen"}
	q := Peo{2, "bike"}

	fmt.Println(getDes(p, q))
	fmt.Println(p.getDes(q))
	fmt.Println(q.getDes(p))

}

func getDes(p Peo, q Peo) ([2]int, [2]string) {
	ids := [2]int{p.id, q.id}
	names := [2]string{p.name, q.name}
	return ids, names
}

//在函数声明时，在其名字之前放上一个变量，即将函数表示为一个方法。
//这个操作会将该函数附加到这种类型上，即相当于为这种类型定义了一个独占的方法。
//比如上面的getDes函数就这个包的，但是下面这个getDes是Peo这个类型的
//而前面小写的这个p是任意的，就如同Java中的this一样，指向哪个实例引用的这个方法
//附加的参数p，叫做方法的接收器（receiver）
func (man Peo) getDes(q Peo) ([2]int, [2]string) {
	ids := [2]int{q.id, q.id}
	names := [2]string{q.name, q.name}
	return ids, names
}

type Peo struct {
	id   int
	name string
}

//接收器参数可以是非指针类型，也可以是指针类型
//当调用一个函数时，会对其每一个参数值进行拷贝，
//如果一个函数需要更新一个变量，
//或者函数的其中一个参数实在太大我们希望能够避免进行这种默认的拷贝，
//这种情况下我们就需要用到指针了。
//对应到我们这里用来更新接收器的对象的方法，
//当这个接受者变量本身比较大时，我们就可以用其指针而不是对象来声明方法
//如个我们设置这个接收器是指针类型
//我们可以通过&的方式去调，也可以不用&的方式去调，
//编译器会隐式地帮我们用&p去调用这个方法。
//这种简写方法只适用于“变量”，包括struct里的字段比如p.X，以及array和slice内的元素比如perim[0]。
//我们不能通过一个无法取到地址的接收器来调用指针方法，比如临时变量的内存地址就无法获取得到
/**
所以总结如下：
1. 不管你的method的receiver是指针类型还是非指针类型，都是可以通过指针/非指针类型进行调用的，编译器会帮你做类型转换。
2. 在声明一个method的receiver该是指针还是非指针类型时，你需要考虑两方面的因素，
第一方面是这个对象本身是不是特别大，如果声明为非指针变量时，调用会产生一次拷贝；
第二方面是如果你用指针类型作为receiver，那么你一定要注意，这种指针类型指向的始终是一块内存地址，就算你对其进行了拷贝，也只是拷贝了这一份地址。
3. 方法理论上也可以用nil指针作为其接收器，尤其当nil对于对象来说是合法的零值时，比如map或者slice
(当你定义一个允许nil作为接收器值的方法的类型时，在类型前面的注释中指出nil变量代表的意义是很有必要的)
*/

//此外，为了避免歧义，在声明方法时，如果一个类型名本身是一个指针的话，是不允许其出现在接收器中的，比如下面这个例子：
//type P *int
//func (P) f() { /* ... */ } // compile error: invalid receiver type

func (man *Peo) getDestiny() {
	fmt.Println(man.name)
	fmt.Println(man.id)
}
func (man Peo) getDestiny2() {
	fmt.Println(man.name)
	fmt.Println(man.id)
}

func main() {

	p := Peo{1, "listen"}
	p.getDestiny()
	p.getDestiny2()
	(&p).getDestiny2()
	(&p).getDestiny()

	pp := &p
	pp.getDestiny2()
	pp.getDestiny()

}
