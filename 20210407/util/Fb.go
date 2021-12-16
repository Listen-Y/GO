package util

func FbNum(num int) int {
	if num <= 0 {
		return 0
	}
	if num == 1 || num == 2 {
		return 1
	}
	return FbNum(num-1) + FbNum(num-2)
}
