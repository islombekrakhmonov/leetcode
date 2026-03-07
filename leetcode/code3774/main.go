package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(absDifference([]int{1, 2}, 1))
}

func absDifference(nums []int, k int) int {
	sort.Ints(nums)

	largestSum := 0
	for i := 1; i <= k; i++ {
		largestSum += nums[len(nums)-i]
	}
	smallestSum := 0
	for i := 0; i < k; i++ {
		smallestSum += nums[i]
	}

	return int(math.Abs(float64(largestSum - smallestSum)))
}
