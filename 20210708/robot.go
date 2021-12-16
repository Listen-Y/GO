package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Println("Welcome to robot talking room...")
	var times = 0
	for {
		_, _ = fmt.Scan(&str)
		str = strings.Replace(str, "?", "！", 1)
		str = strings.Replace(str, "？", "！", 1)
		str = strings.Replace(str, "吗", "", 1)
		fmt.Println(str)
		times++
		if times == 10 {
			break
		}
	}
	fmt.Println("bye...")
}
