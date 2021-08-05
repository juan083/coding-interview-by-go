package main

import (
	"fmt"
	"math"
)

/**
二分查找法：
mid: 中间位置，比较值(mid 与 r 对应的值做比较)：
mid < r：说明最小值在[left, mid], right = mid
mid == r：说明最小值在[left, right - 1], right--
mid > r：说明最小值在(mid, right], left = mid + 1
*/
func minArray(numbers []int) int {
	length := len(numbers)
	if length == 0 {
		return -1
	}
	l, r := 0, length-1
	for l < r {
		mid := int(math.Ceil(float64((r-l)/2))) + l
		if numbers[mid] > numbers[r] {
			l = mid + 1
		} else if numbers[mid] < numbers[r] {
			r = mid
		} else {
			r--
		}
	}
	return numbers[l]
}

func main() {
	numbers := []int{3, 4, 5, 1, 2}
	index := minArray(numbers)
	fmt.Println(index)
}
