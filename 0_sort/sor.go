package main

// 冒泡
func BubbleSort(arr []int) []int {
	if arr == nil || len(arr) < 1 {
		return arr
	}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

// 快排
func QuickSort(arr []int) []int {
	if arr == nil || len(arr) < 1 {
		return arr
	}
	base := arr[0]
	low := make([]int, 0)
	high := make([]int, 0)
	mid := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		if arr[i] < base {
			low = append(low, arr[i])
		} else if arr[i] > base {
			high = append(high, arr[i])
		} else {
			mid = append(mid, arr[i])
		}
	}
	low, high = QuickSort(low), QuickSort(high)
	result := append(append(low, mid...), high...)
	return result
}

func HeapSortMax(arr []int, length int) []int {
	if length <= 1 {
		return arr
	}
	depth := length/2 - 1 //二叉树深度
	for i := depth; i >= 0; i-- {
		topmax := i //假定最大的位置就在i的位置
		leftchild := 2*i + 1
		rightchild := 2*i + 2
		if leftchild <= length-1 && arr[leftchild] > arr[topmax] { //防止越过界限
			topmax = leftchild
		}
		if rightchild <= length-1 && arr[rightchild] > arr[topmax] { //防止越过界限
			topmax = rightchild
		}
		if topmax != i {
			arr[i], arr[topmax] = arr[topmax], arr[i]
		}
	}
	return arr
}

func HeapSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		lastlen := length - i
		HeapSortMax(arr, lastlen)
		if i < length {
			arr[0], arr[lastlen-1] = arr[lastlen-1], arr[0]
		}
	}
	return arr
}
