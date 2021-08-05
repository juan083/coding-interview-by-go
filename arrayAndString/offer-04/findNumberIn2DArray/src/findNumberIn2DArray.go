package main

import "fmt"

func findNumberIn2DArray(matrix [][]int, target int) bool {
	xLen := len(matrix)
	// 如果数组没有元素，则直接返回false
	if xLen == 0 {
		return false
	}

	// x是行坐标，y是列坐标
	x, y := 0, len(matrix[0])-1
	for x < xLen && y >= 0 {
		if matrix[x][y] == target {
			return true
		} else if matrix[x][y] > target {
			// 当前元素 > target，则说明target在上部分
			y--
		} else {
			// 当前元素 < target，则说明target在右部分
			x++
		}
	}
	return false
}

func main() {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	flag := findNumberIn2DArray(matrix, 15)
	fmt.Println(flag)
}
