package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

func main1() {
	fmt.Println("flags: ", os.Args)

	// 解析数据
	var f string
	var p int
	var u string

	flag.StringVar(&f, "f", "", "文件路径,默认为空")
	flag.IntVar(&p, "p", 0, "端口号,默认为0")
	flag.StringVar(&u, "f", "root", "用户,默认为root")
	flag.Parse() // 必须进行这步,进行数据解析
}

func main2() {
	// 演示将map进行序列化
	data := make(map[string]interface{}, 0)
	data["name"] = "listen"
	data["age"] = 22
	data["id"] = 329043988827384909
	data["score"] = 98.5

	marshal, err := json.Marshal(data)
	if err != nil {
		return
	}
	fmt.Println(string(marshal))
}

func main3() {
	jsonData := "{\"age\":22,\"id\":329043988827384909,\"name\":\"listen\",\"score\":98.5}"

	var data map[string]interface{}

	// 这里的map不需要make也能使用，因为在json.Unmarshal中已经进行了make
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		return
	}
	fmt.Println(data)
}

type Stu struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

func main() {
	stu := Stu{
		Name: "Listen",
		Id:   1001,
	}

	marshal, err := json.Marshal(stu)
	if err != nil {
		return
	}
	fmt.Println(string(marshal))
}
