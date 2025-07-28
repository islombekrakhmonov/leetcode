package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxAscendingSum([]int{10, 20, 30, 5, 10, 50}))
}

func maxAscendingSum(nums []int) int {

	var maxSum = math.MinInt

	for startIndex := 0; startIndex < len(nums); startIndex++ {
		for endIndex := startIndex + 1; endIndex <= len(nums); endIndex++ {
			isAscending, sum := isAscending(nums[startIndex:endIndex])
			if isAscending && sum > maxSum {
				maxSum = sum
			}
		}
	}

	return maxSum
}

func isAscending(nums []int) (bool, int) {
	sum := 0

	if len(nums) == 1 {
		return true, nums[0]
	}

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] >= nums[i+1] {
			return false, 0
		}
		sum += nums[i]
	}

	sum += nums[len(nums)-1]

	return true, sum
}
