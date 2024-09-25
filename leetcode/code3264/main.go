package main

import (
	"fmt"
)

func main() {
	fmt.Println(getFinalState([]int{2, 1, 3, 5, 6}, 5, 2))
}

func getFinalState(nums []int, k int, multiplier int) []int {
	for k > 0 {
		min := 0
		for i := range nums {
			if nums[min] > nums[i] {
				min = i
			}
		}
		nums[min] *= multiplier
		k--
	}

	return nums
}
