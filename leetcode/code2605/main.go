package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(minNumber([]int{7, 5, 6}, []int{1, 4}))
}

func minNumber(nums1 []int, nums2 []int) int {
	var existingNums []int
	var min1, min2 = math.MaxInt, math.MaxInt

	for _, num1 := range nums1 {
		if num1 < min1 {
			min1 = num1
		}
		for _, num2 := range nums2 {
			if num2 < min2 {
				min2 = num2
			}
			if num1 == num2 {
				existingNums = append(existingNums, num1)
			}
		}
	}

	sort.Ints(existingNums)

	if len(existingNums) > 0 {
		return existingNums[0]
	}

	minCombined := min1*10 + min2
	if minCombined < min2*10+min1 {
		return minCombined
	}

	return min2*10 + min1
}
