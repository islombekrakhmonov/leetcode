package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minimumOperations([]int{2}))
}

func minimumOperations(nums []int) int {
	l:= len(nums)
	if l == 1 && nums[0] == 0 {
		return 0
	}

	operations := 0
	for {
		min := math.MaxInt
		hasNonZero := false

		for i := 0; i < l; i++ {
			if nums[i] > 0 {
				hasNonZero = true
				if min > nums[i] {
					min = nums[i]
				}
			}
		}

		if !hasNonZero {
			break
		}

		for i := 0; i < l; i++ {
			if nums[i] > 0 {
				nums[i] -= min
			}
		}

		operations++
	}

	return operations
}