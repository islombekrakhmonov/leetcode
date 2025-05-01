package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minimumPairRemoval([]int{2, 2, -1, 3, -2, 2, 1, 1, 1, 0, -1}))
}

func minimumPairRemoval(nums []int) int {
	var count int

	for !isIncreasing(nums) {
		minimumSumPairIndex := make(map[int]int)
		var minSum = math.MaxInt
		var minIndex int

		for i := 0; i < len(nums)-1; i++ {
			sum := nums[i] + nums[i+1]
			if sum < minSum {
				minSum = sum
				minIndex = i
			}
		}

		if minimumSumPairIndex[minSum] < len(nums)-1 {
			nums = append(nums[:minIndex], append([]int{minSum}, nums[minIndex+2:]...)...)
		}
		count++
	}

	return count
}

func isIncreasing(nums []int) bool {

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return false
		}
	}

	return true
}
