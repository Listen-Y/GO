package main

import (
	"container/heap"
	"container/list"
	"container/ring"
	"fmt"
	"sort"
	"strconv"
)

type aStructure struct {
	person string
	height int
	weight int
}

func main1() {
	mySlice := make([]aStructure, 0)
	mySlice = append(mySlice, aStructure{"bike", 1, 2},
		aStructure{"faker", 100, 200},
		aStructure{"listen", 10, 20}) // 定义切片

	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].height < mySlice[j].height
	}) // 按照身高由小到大排序
	fmt.Println("<:", mySlice)

	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].height > mySlice[j].height
	}) // 按照身高由大到小排序
	fmt.Println(">:", mySlice)
}

func main2() {
	// 根据参数值（10），计算出合适的容量。
	// 一个map 中会包含很多桶，每个桶中可以存放8个键值对。
	info := make(map[string]string, 10)
	//info[nil] = "a"
	// info["n4"] = nil // 报错
	//fmt.Println(info[nil])

	info["n1"] = "武沛齐"
	info["n2"] = "alex"

	info["n3"] = "aaa"
	info["n1"] = "bbb"
	fmt.Println(info["n1"])

}

type student struct {
	name string
	age  int32
	id   int64
}

type heapStudents []student

func (h *heapStudents) Push(x interface{}) {
	val, err := x.(student)
	if err != true {
		fmt.Println(fmt.Errorf("value error"))
	}
	*h = append(*h, val) // 每次push都会新开辟一个切片, 如果需要这里可以做一个优化
	heap.Init(h)
}

func (h *heapStudents) Pop() interface{} {
	old := *h
	val := old[0]
	*h = old[1:] // 每次pop都会新开辟一个切片, 如果需要这里可以做一个优化
	heap.Init(h)
	return val
}

func (h heapStudents) Len() int {
	return len(h)
}

func (h heapStudents) Less(i, j int) bool {
	//return h[i].age < h[j].age // 小顶堆, 用于升序排序
	return h[i].age > h[j].age // 大顶堆, 用于降序排序
}

func (h heapStudents) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func main4() {
	myHeap := heapStudents{}
	myHeap = append(myHeap,
		student{"a", 6, 11},
		student{"b", 1, 11},
		student{"c", 9, 11},
		student{"d", 3, 11},
		student{"e", 5, 11},
		student{"f", 7, 11})
	fmt.Println(myHeap)

	fmt.Println("heap init...")
	myHeapInit := &myHeap
	heap.Init(myHeapInit)
	fmt.Println(myHeap)

	fmt.Println("push...")
	myHeap.Push(student{"g", 0, 11})
	myHeap.Push(student{"h", 2, 11})
	fmt.Println(myHeap)

	fmt.Println("pop...")
	newHeap := make(heapStudents, len(myHeap))
	copy(newHeap, myHeap)
	for i := 0; i < len(myHeap); i++ {
		fmt.Println(newHeap.Pop().(student))
	}
}

type temp struct {
	//kvs map[string]string
	id int64
}

func main5() {
	// values := list.New()
	s1 := student{"a", 1, 2}
	s2 := student{
		name: "a",
		age:  1,
		id:   2,
	}
	fmt.Println(s1 == s2)

	t1 := temp{
		//kvs: make(map[string]string, 0),
		id: 0,
	}
	t2 := temp{
		//kvs: make(map[string]string, 0),
		id: 0,
	}
	//t1.kvs["a"] = "a"
	//t1.kvs["a"] = "a"

	fmt.Println(t1 == t2) // map不支持比较, 如果结构体里有map,那么这个结构体也不支持比较
}

func main6() {
	values := list.New()
	e1 := values.PushBack("Two")   // 链表尾部插入
	e2 := values.PushBack("Three") //
	values.PushFront("Zero")       // 链表头部插入
	values.InsertBefore("One", e1) // 指定位置插入
	values.InsertAfter("Four", e2)
	values.Remove(e2) // 删除节点
	values.Remove(e2)
	values.InsertAfter("FiveFive", e2) // 此时不会插入到values链表
	values.PushBackList(values)        // 链表整个插入

	printList(values) // 输出链表

	values.Init() // 清空链表

	fmt.Printf("After Init(): %v\n", values)
	for i := 0; i < 20; i++ {
		values.PushFront(strconv.Itoa(i)) // 因为init清空了链表,所以此时会像新建一个链表一样是一个新的链表
	}
	printList(values)

}

func printList(l *list.List) {
	fmt.Println("倒着输出....")
	for t := l.Back(); t != nil; t = t.Prev() {
		fmt.Print(t.Value, " ")
	}
	fmt.Println()
	fmt.Println("正着输出....")
	for t := l.Front(); t != nil; t = t.Next() {
		fmt.Print(t.Value, " ")
	}
	fmt.Println()
	fmt.Println()
}

func main() {
	size := 10
	myRing := ring.New(size)
	fmt.Println("Empty ring:", *myRing)
	for i := 0; i < myRing.Len(); i++ {
		myRing.Value = i
		myRing = myRing.Next() // 这俩步操作是在添加数据
	}
	printRing(myRing)
	myRing = myRing.Move(-1) // 向前移动一位
	printRing(myRing)

	sum := 0
	myRing.Do(func(x interface{}) {
		t := x.(int)
		sum += t
	})
	fmt.Println("Sum:", sum)
}

func printRing(r *ring.Ring) {
	fmt.Printf("len: %d\n", r.Len())
	for i := 0; i < r.Len(); i++ {
		fmt.Print(r.Value, " ")
		r = r.Next()
	}
	fmt.Println()
}
