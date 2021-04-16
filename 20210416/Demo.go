package main

import (
	"fmt"
	"sort"
)

/*
map学习
*/
func main1() {

	map1 := make(map[string]int)
	map1["a"] = 1
	map1["b"] = 2
	map1["c"] = 3

	map2 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	delete(map1, "b")
	map1["a"] = map1["a"] + 1

	show(map1, map2)

}

func show(map1 map[string]int, map2 map[string]int) {
	for k, v := range map1 {
		fmt.Printf("k = %s, v = %d\n", k, v)
	}
	fmt.Println("-------------------------")
	for k, v := range map2 {
		fmt.Printf("k = %s, v = %d\n", k, v)
	}
	fmt.Println("*************************")
}

func main() {
	map1 := make(map[string]int)
	map1["d"] = 1
	map1["b"] = 2
	map1["c"] = 3

	//每次往slice中方数据的时候都会给len位置上放，所以初始化slice的时候应该len设置位0，cap设置成最优
	names := make([]string, 0, 4)
	fmt.Printf("names: len =%d, cap = %d\n", len(names), cap(names))
	for name := range map1 {
		names = append(names, name)
		fmt.Printf("names: len =%d, cap = %d\n", len(names), cap(names))
	}
	//sort这个工具类用于排序
	sort.Strings(names)
	fmt.Println(names)

}
