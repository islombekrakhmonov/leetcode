package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(minimumSum([]int{8, 6, 1, 5, 3}))
}

func minimumSum(nums []int) int {

	l := len(nums)

	minSumOfTriplets := math.MaxInt

	for i := 0; i < l-2; i++ {
		for j := i + 1; j < l-1; j++ {
			for k := j + 1; k < l; k++ {
				if nums[i] < nums[j] && nums[k] < nums[j] {
					sum := nums[i] + nums[j] + nums[k]
					if sum < minSumOfTriplets {
						minSumOfTriplets = sum
					}
				}
			}
		}
	}

	if minSumOfTriplets != math.MaxInt {
		return minSumOfTriplets
	}

	return -1
}
