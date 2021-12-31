package main

import (
	"fmt"
	"sort"
	"strconv"
	"time"
)

type USB interface {
	// 声明方法
	Start(str string) string
	Stop(str string) string
}
type Phone struct {
	Id int
}

func (p *Phone) Start(str string) string {
	// 将int转为string, 这种能避免乱码
	return strconv.Itoa(p.Id) + " Phone start: " + str
}
func (p *Phone) Stop(str string) string {
	return strconv.Itoa(p.Id) + " Phone stop: " + str
}

type Computer struct {
	Id int
}

func (c *Computer) Start(str string) string {
	return strconv.Itoa(c.Id) + " Computer start: " + str
}
func (c *Computer) Stop(str string) string {
	return strconv.Itoa(c.Id) + " Computer stop: " + str
}

func Working(usb USB, desc string) {
	fmt.Println(usb.Start(desc))
	time.Sleep(time.Second)
	fmt.Println(usb.Stop(desc))
}

func main6() {
	usbs := [3]USB{}
	usbs[0] = &Computer{}
	usbs[1] = &Phone{}
	usbs[2] = &Computer{}

	fmt.Println(usbs)
}

func main1() {
	phone := Phone{
		Id: 0,
	}

	computer := Computer{
		Id: 1001,
	}

	Working(&phone, "手机") // 接口的本质是指针, 所以需要取地址
	// 方式二
	var usb USB
	usb = &computer
	Working(usb, "电脑")
}

type Stu struct {
	Id int
}

type Stus []Stu

func (s Stus) Len() int { // 切片是引用传递,所以给其数据类型实现sort接口
	return len(s) // 这里需要注意
}

func (s Stus) Less(i, j int) bool {
	return s[i].Id > s[j].Id
}

func (s Stus) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main2() {

	stus := Stus{}
	stus = append(stus, Stu{
		Id: 0,
	}, Stu{
		Id: 1,
	}, Stu{
		Id: 2,
	}, Stu{
		Id: 3,
	}, Stu{
		Id: 4,
	}, Stu{
		Id: 5,
	}, Stu{
		Id: 6,
	}, Stu{
		Id: 7,
	})
	sort.Sort(stus)
	fmt.Println(stus)
}

type Integer int

func (i Integer) Start(str string) string {
	return fmt.Sprint(i) + " Integer start: " + str
}
func (i Integer) Stop(str string) string {
	return fmt.Sprint(i) + " Integer stop: " + str
}

func Test(usb USB) {
	if usb == nil {
		return
	}
	integer := usb.(Integer)

	fmt.Println(integer)
}

func main4() {
	var i Integer
	i = 10
	fmt.Println(i.Stop("desc"))

	Test(i)
	fmt.Println(i)
}

type T interface {
}

func main5() {
	var t T

	t = Stu{
		Id: 0,
	}
	fmt.Println(t)

	t = Integer(1)
	fmt.Println(t)

	t = Computer{
		Id: 0,
	}
	fmt.Println(t)
}

type Point struct {
	x int
	y int
}

func main7() {
	var t interface{}
	point := Point{
		x: 0,
		y: 0,
	}
	t = point
	fmt.Println(t)

	// 如何将t重新赋予给Point
	var point1 Point
	point1 = t.(Point)
	fmt.Println(point1.x, point1.y)
}

func main8() {
	var t interface{}

	var num float64
	num = 1.1
	t = num

	num2, ok := t.(float64)
	if !ok {
		fmt.Println("convert error")
	} else {
		fmt.Printf("t type: %T value: %v\n", num2, num2)
	}
}

func typeJudge(items ...interface{}) {
	for index, v := range items {
		switch v.(type) {
		case int:
			fmt.Printf("第%v个参数 类型是%T 值是%v\n", index, v.(int), v)
		case float32:
			fmt.Printf("第%v个参数 类型是%T 值是%v\n", index, v.(float32), v)
		case string:
			fmt.Printf("第%v个参数 类型是%T 值是%v\n", index, v.(string), v)
		case Point:
			fmt.Printf("第%v个参数 类型是%T 值是%v\n", index, v.(Point), v)
		case *Point:
			fmt.Printf("第%v个参数 类型是%T 值是%v\n", index, v.(*Point), v) // 注意这个也是可以的
		case int64:
			fmt.Printf("第%v个参数 类型是%T 值是%v\n", index, v.(int64), v)
		default:
			fmt.Printf("第%v个参数 类型是%T 值是%v\n", index, v, v)
		}
	}
}

func main() {
	itmes := make([]interface{}, 0)
	var num int = 1
	itmes = append(itmes, num)
	var num2 float32 = 1.1
	itmes = append(itmes, num2)
	itmes = append(itmes, "listen")
	itmes = append(itmes, Point{
		x: 0,
		y: 0,
	})
	itmes = append(itmes, &Point{
		x: 1,
		y: 1,
	})
	var num3 int64 = 9
	itmes = append(itmes, num3)

	typeJudge(itmes...)
}
