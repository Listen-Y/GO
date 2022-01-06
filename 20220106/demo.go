package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
	Id   int
}

func main() {
	user := User{
		Name: "",
		Age:  0,
		Id:   0,
	}

	typeOf := reflect.TypeOf(user)

	// 这里返回的是一个newUser 的指针
	newUser := reflect.New(typeOf)
	// 获取newUser 的一个指针
	elem := newUser.Elem()
	elem.FieldByName("Name").SetString("listen")
	elem.FieldByName("Age").SetInt(20)
	elem.FieldByName("Id").SetInt(1001)

	u := newUser.Interface().(*User)
	fmt.Println(u)
}

func main1() {
	user := User{
		Name: "",
		Age:  0,
		Id:   0,
	}

	// 获取*user的value
	valueOf := reflect.ValueOf(&user)

	// 获取 *user的一个新的指针
	elem := valueOf.Elem()
	elem.FieldByName("Name").SetString("listen")
	elem.FieldByName("Age").SetInt(20)
	elem.FieldByName("Id").SetInt(1001)

	fmt.Println(user)
}
