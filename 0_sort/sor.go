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

// 堆排序
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

func HeapSortMax(arr []int, length int) {
	if length <= 1 {
		return
	}
	depth := length/2 - 1
	for i := depth; i >= 0; i-- {
		maxIndex := i
		leftIndex := i*2 + 1
		rightIndex := i*2 + 2
		if leftIndex <= length-1 && arr[leftIndex] > arr[maxIndex] {
			maxIndex = leftIndex
		}
		if rightIndex <= length-1 && arr[rightIndex] > arr[maxIndex] {
			maxIndex = rightIndex
		}
		if maxIndex != i {
			arr[i], arr[maxIndex] = arr[maxIndex], arr[i]
		}
	}
}

// 二分查找
func binSearch(arr []int, target int) bool {
	low := 0
	hight := len(arr) - 1
	for low <= hight {
		mid := (low + hight) / 2
		if arr[mid] > target {
			hight = mid - 1
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			return true
		}
	}
	return false
}

// 选择排序
