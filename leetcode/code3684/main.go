package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxKDistinct([]int{1, 1, 1, 2, 2, 2}, 6))
}

func maxKDistinct(nums []int, k int) []int {
	var output []int

	sort.Ints(nums)

	for i := len(nums) - 1; i >= 0; i-- {
		if i != 0 {
			if nums[i] == nums[i-1] {
				continue
			}
		}
		output = append(output, nums[i])
		if len(output) == k {
			break
		}
	}

	return output
}
