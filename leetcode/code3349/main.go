package main

import (
	"fmt"
)

func main() {
	fmt.Println(hasIncreasingSubarrays([]int{-15, -13, 4, 7}, 2))
}

func hasIncreasingSubarrays(nums []int, k int) bool {
	subArrayIndexes := []int{}
	for startIndex := 0; startIndex < len(nums); startIndex++ {
		for endIndex := startIndex + 1; endIndex <= len(nums); endIndex++ {
			if endIndex-startIndex == k {
				subarray := nums[startIndex:endIndex]
				if isIncreasing(subarray) {
					subArrayIndexes = append(subArrayIndexes, startIndex)
				}
			}
		}
	}

	for i := 0; i < len(subArrayIndexes)-1; i++ {
		for j := i + 1; j < len(subArrayIndexes); j++ {
			if subArrayIndexes[j]-subArrayIndexes[i] == k {
				return true
			}
		}
	}

	return false
}

func isIncreasing(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] >= nums[i+1] {
			return false
		}
	}

	return true
}
