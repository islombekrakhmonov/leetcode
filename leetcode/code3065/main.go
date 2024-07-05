package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minOperations([]int{2, 11, 10, 1, 3}, 10))
}

func minOperations(nums []int, k int) int {

	sort.Ints(nums)

	var (
		output int
	)

	for i := 0; i < len(nums); i++ {
		if nums[i] < k {
			output++
		} else {
			return output
		}
	}

	return output
}
