package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minSubsequence([]int{4, 3, 10, 9, 8}))
}

func minSubsequence(nums []int) []int {
	var (
		output []int
		sum    int
	)

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	for _, num := range nums {
		sum += num
	}

	currentSum := 0
	for i := 0; i < len(nums); i++ {
		currentSum += nums[i]
		output = append(output, nums[i])
		if currentSum > sum-currentSum {
			break
		}
	}

	return output
}

// func sum(nums []int) int {
// 	sum := 0
// 	for _, num := range nums {
// 		sum += num
// 	}
// 	return sum
// }
