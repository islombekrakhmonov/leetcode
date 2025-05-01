package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestMonotonicSubarray([]int{1, 4, 3, 3, 2}))
}

func longestMonotonicSubarray(nums []int) int {
	var maxLength int

	for startIndex := 0; startIndex < len(nums); startIndex++ {
		for endIndex := startIndex + 1; endIndex <= len(nums); endIndex++ {
			subArray := nums[startIndex:endIndex]
			if isStrictlyDecreasing(subArray) || isStrictlyIncreasing(subArray) {
				fmt.Println(subArray)
				if len(subArray) > maxLength {
					maxLength = len(subArray)
				}
			}
		}
	}

	return maxLength
}

func isStrictlyIncreasing(array []int) bool {
	for i := 0; i < len(array)-1; i++ {
		if array[i] >= array[i+1] {
			return false
		}
	}

	return true
}

func isStrictlyDecreasing(array []int) bool {
	for i := 0; i < len(array)-1; i++ {
		if array[i] <= array[i+1] {
			return false
		}
	}

	return true
}
