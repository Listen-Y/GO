package main

import (
	"fmt"
	"sort"
)

/*
定义额外的切片type类型，让其实现Len，Less，Swap方法，即可让其实现sort接口
*/
type StringSlices []string

func (names StringSlices) Len() int {
	return len(names)
}

func (names StringSlices) Less(i, j int) bool {
	return len(names[i]) > len(names[j])
}

func (names StringSlices) Swap(i, j int) {
	names[i], names[j] = names[j], names[i]
}

func mainB() {
	names := StringSlices{"a", "aa", "aaa", "aaaa", "aaaaa"}
	sort.Sort(names)
	fmt.Println(names)
}

/*
自定义结构体，实现自定义排序，让其按照id进行排序
*/
type Person struct {
	Id    int
	Name  string
	Class string
}

type Persons []*Person //这是一个数组指针，是一个数组里面存放的是指针

func (persons Persons) Len() int {
	return len(persons)
}

func (persons Persons) Less(i, j int) bool {
	return (persons[i]).Id > (persons[j]).Id
}

func (persons Persons) Swap(i, j int) {
	persons[i], persons[j] = persons[j], persons[i]
}

func main() {
	persons := Persons{
		{1, "a", "b1"},
		{2, "aa", "b2"},
		{3, "aaa", "b3"},
		{4, "aaaa", "b4"},
		{5, "aaaaa", "b5"},
		{6, "aaaaaa", "b6"},
	}
	sort.Sort(persons)
	for _, val := range persons {
		fmt.Println(*val)
	}
}

func mainA() {

	testNames := []string{"test", "", "test", "testing", "testT"}
	namesCache := make(map[string]bool)

	for _, val := range testNames {
		if val == "" {
			continue
		}
		if namesCache[val] {
			continue
		}
		namesCache[val] = true
	}

	names := make([]string, len(namesCache))
	for k := range namesCache {
		names = append(names, k)
	}
	fmt.Println(names)

}
