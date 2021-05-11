package main

import (
	"fmt"
	"reflect"
)

/*
反射其实就是为了能够在运行期间获取并到数据的类型和数据的变量和数据的方法，
并且运行时更新变量和检查它们的值、调用它们的方法和它们支持的内在操作
 reflect.TypeOf 接受任意的 interface{} 类型，并以 reflect.Type 形式返回其动态类型
fmt.Printf 提供了一个缩写 %T 参数，内部使用 reflect.TypeOf 来输出
*/
func main5() {
	t := reflect.TypeOf(3)  // a reflect.Type
	fmt.Println(t.String()) // "int"
	fmt.Println(t)          // "int"
}

/*
reflect.ValueOf 接受任意的 interface{} 类型
reflect.ValueOf 返回的结果也是具体的类型，但是 reflect.Value 也可以持有一个接口值。
reflect.Type 类似，reflect.Value 也满足 fmt.Stringer 接口，
但是除非 Value 持有的是字符串，否则 String 方法只返回其类型。而
使用 fmt 包的 %v 标志参数会对 reflect.Values 特殊处理
*/

func main6() {
	v := reflect.ValueOf(3) // a reflect.Value
	fmt.Println(v)          // "3"
	fmt.Printf("%v\n", v)   // "3"
	fmt.Println(v.String()) // NOTE: "<int Value>"
}

/*
reflect.ValueOf 的逆操作是 reflect.Value.Interface 方法。
它返回一个 interface{} 类型，装载着与 reflect.Value
*/

func main7() {
	x := reflect.ValueOf(2)
	v := x.Interface()
	data, _ := v.(int)
	fmt.Println(data)
}

/*
.Kind也可以直接获取其type
*/
func main8() {
	v := reflect.ValueOf("aa")
	fmt.Println(v.Kind() == reflect.String)
	vv := reflect.TypeOf("aa")
	fmt.Println(vv.String() == "string")
}

/*
通过reflect.ValueOf(x)返回的reflect.Value都是不可取地址的,因为获得的值仅仅是参数的拷贝副本
但是通过.Elem获得的是可以取地址的,因为它是以引用方式生成的，指向另一个变量
可以通过调用reflect.Value的CanAddr方法来判断其是否可以被取地址
*/

func main9() {
	x := 2                   // value   type    variable?
	a := reflect.ValueOf(2)  // 2       int     no
	b := reflect.ValueOf(x)  // 2       int     no
	c := reflect.ValueOf(&x) // &x      *int    no
	d := c.Elem()            // 2       int     yes (x)
	fmt.Println(a.CanAddr())
	fmt.Println(b.CanAddr())
	fmt.Println(c.CanAddr())
	fmt.Println(d.CanAddr()) //全为FALSE，就这个为true
}

/*
通过反射改变值
对一个不可取地址的reflect.Value调用Set方法也会导致panic异常
比如Json.Unmarshal函数使用了下面的反射机制类修改数据变量的每个成员
*/
func main() {
	//方式一
	num := 10
	v := reflect.ValueOf(&num).Elem()
	p := v.Addr().Interface().(*int)
	*p = 11
	fmt.Println(num)

	//方式二
	val := 10
	reflect.ValueOf(&val).Elem().Set(reflect.ValueOf(11)) //直接通过Value的内置方法改变数据
	fmt.Println(val)
}
