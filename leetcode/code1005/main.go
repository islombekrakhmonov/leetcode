package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(largestSumAfterKNegations([]int{-2, 9, 9, 8, 4}, 5))
}

func largestSumAfterKNegations(nums []int, k int) int {

	fmt.Println(nums)

	for k != 0 {
		sort.Slice(nums, func(i, j int) bool {
			return nums[i] < nums[j]
		})
		for i := 0; i < len(nums); i++ {
			if k != 0 {
				if nums[i] <= 0 {
					nums[i] = -nums[i]
					k--
				}
			} else {
				break
			}
		}
		sort.Slice(nums, func(i, j int) bool {
			return nums[i] < nums[j]
		})
		if k != 0 {
			nums[0] = -nums[0]
			k--
		}
	}

	sum := 0

	for _, v := range nums {
		sum += v
	}

	return sum
}

//Sort the array in non decreasing order.
// All the smallest numbers/negative numbers will be placed at the beginning , so start flippin them from starting and after each flip decrement the K.
// If there are still K remaining and all negatives have been flipped , then flip the smallest absolute value in the array for remaining k times.
