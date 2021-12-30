package main

import "fmt"

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

func main() {

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
