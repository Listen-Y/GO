package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func randSort(arr []int) {
	rand.Seed(time.Now().Unix())
	indexs := rand.Perm(len(arr))
	// 交换位置
	for i, i2 := range indexs {
		arr[i], arr[i2] = arr[i2], arr[i]
	}
}

func main() {
	arr := []int{33, 99, 66, 44, 22}
	count := 0
	startTime := time.Now()
	for {
		if sort.SliceIsSorted(arr, func(i, j int) bool {
			return arr[i] < arr[j]
		}) {
			fmt.Printf("有序了 arr:%v，一共排序了%v次\n", arr, count)
			break
		} else {
			randSort(arr)
			count++
		}
	}
	fmt.Printf("耗时：%v\n", time.Now().Sub(startTime))
}
