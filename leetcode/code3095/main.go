package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minimumSubarrayLength([]int{2, 1, 8}, 10))
}

func minimumSubarrayLength(nums []int, k int) int {

	minimumLength := math.MaxInt

	for startIndex := 0; startIndex < len(nums); startIndex++ {
		for endIndex := startIndex + 1; endIndex <= len(nums); endIndex++ {
			result := 0
			subarray := nums[startIndex:endIndex]
			for _, v := range subarray {
				result |= v
			}
			if result >= k {
				if len(subarray) < minimumLength {
					minimumLength = len(subarray)
				}
			}

		}
	}

	if minimumLength == math.MaxInt {
		return -1
	}

	return minimumLength
}
