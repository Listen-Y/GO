package main

import (
	"encoding/json"
	"fmt"
)

func main1() {
	var kv map[string]string
	kv = make(map[string]string, 4)

	kv["a"] = "aa"
	kv["b"] = "bb"

	val, find := kv["b"]
	if find { // 是否存在
		fmt.Println(val)
	}

	// 遍历
	for k, v := range kv {
		fmt.Println("k:", k, "v:", v)
	}

	fmt.Println("------------------------------------")

	// 复杂map的遍历
	infos := make(map[int]map[string]string, 4)
	_, ok := infos[1]
	if !ok { // 不存在
		infos[1] = make(map[string]string, 2)
	}
	values := infos[1] // 注意这里要获取make后的map
	values["name"] = "listen"
	values["age"] = "22"
	values["id"] = "001"
	//_ ,ok = infos[2]
	if infos[2] == nil { // 不存在
		infos[2] = make(map[string]string, 2)
	}
	values = infos[2]
	values["name"] = "tom"
	values["age"] = "21"
	values["id"] = "002"

	for _, v := range infos {
		for kk, vv := range v {
			fmt.Println(kk, ":", vv)
		}
	}

}

func main2() {

	// 申明一个map切片
	var stu []map[string]string
	stu = make([]map[string]string, 2) // 此时意思是切片里有俩个map引用, 所以也只能存储俩个map

	if stu[0] == nil {
		stu[0] = make(map[string]string, 2)
		stu[0]["name"] = "tom"
		stu[0]["age"] = "22"
	}

	if stu[1] == nil {
		stu[1] = make(map[string]string, 2)
		stu[1]["name"] = "listen"
		stu[1]["age"] = "21"
	}

	// 但是注意这里有个坑, 如果使用stu[2]就会出现panic, 这里要注意, 所以如果需要再次添加必须使用append
	s := map[string]string{
		"name": "lisa",
		"age":  "20",
	}
	stu = append(stu, s)

	fmt.Println(stu)
}

type Person struct {
	Name  string `json:"name"`
	age   int    `json:"age"`
	infos map[string]string
}

func main3() {
	p1 := Person{
		Name:  "Listen",
		age:   22,
		infos: make(map[string]string, 2), // 必须
	}
	fmt.Println(p1)

	var p2 = new(Person) // 一个Person的指针
	(*p2).Name = "tom"
	p2.age = 21 // 这俩种方式都支持
	(*p2).infos = make(map[string]string, 2)
	fmt.Println(*p2)
}

type A struct {
	Name string
}

type B struct {
	Name string
}

func main4() {
	a := A{
		Name: "",
	}
	b := B{
		Name: "",
	}

	b = B(a)
	fmt.Println(b)
}

type Cat struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (c *Cat) eat() {
	c.Name = "listen"
	fmt.Println(c.Name, ": eating")
}

func GetCat(name string, age int) *Cat {
	//if ...
	return &Cat{
		Name: "",
		Age:  0,
	}
}

func main6() {
	cat := Cat{
		Name: "bike",
		Age:  0,
	}
	cat.eat()

	fmt.Println(cat)
	test(&cat)
	fmt.Println(cat)
}

func test(cat *Cat) {
	cat.Name = "a"
}

func main5() {

	cat := Cat{
		Name: "bike",
		Age:  0,
	}
	marshal, _ := json.Marshal(cat)
	fmt.Println(string(marshal))
}

type People struct {
	Name string
	Age  int
	Id   int
}

func (p *People) showInfo() {
	fmt.Println("info: ", p)
}

type Man struct {
	People
	Sex string
}

func (m *Man) showManInfo() {
	fmt.Println("man info: ", m)
}

func (m *Man) showInfo() {
	fmt.Println("man info: ", m)
}

func main7() {
	man := Man{
		People: People{
			Name: "Listen",
			Age:  1,
			Id:   1,
		},
		Sex: "男",
	}

	man.showManInfo()
	man.showInfo()
	// 或者
	man.People.showInfo()

}

type Q struct {
	Name string
}

type E struct {
	Name string
}

type W struct {
	Q
	E
	//Name string
}

func main8() {
	w := W{
		Q: Q{
			Name: "Q",
		},
		E: E{
			Name: "E",
		},
		//Name: "W",
	}
	//fmt.Println(w.Name)// error
	fmt.Println(w.Q.Name) // right
}

type R struct {
	Name string
}

type T struct {
	*R
}

func main9() {
	t := T{
		R: &R{
			Name: "",
		},
	}
	fmt.Println(t)

}

type Y struct {
	int
	string
}

func main() {
	y := Y{
		int:    0,
		string: "",
	}
	fmt.Println(y.string)
	fmt.Println(y.int)
}
