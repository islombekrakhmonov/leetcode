package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumDifference([]int{20, 200,300, 1000}, 3))
}

func minimumDifference(nums []int, k int) int {
	if k <= 1 {
        return 0
    }

	sort.Ints(nums)
    minDiff := nums[k-1] - nums[0]
	fmt.Println(minDiff)
    for i := 1; i <= len(nums)-k; i++ {
        diff := nums[i+k-1] - nums[i]
        if diff < minDiff {
            minDiff = diff
        }
    }

	return minDiff
}
