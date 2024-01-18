package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumProduct([]int{1, 2, 3, 4}))
}

func maximumProduct(nums []int) int {
	sort.Ints(nums)

	a := (nums[len(nums)-1] * nums[len(nums)-2] * nums[len(nums)-3])
	b := (nums[0] * nums[1] * nums[len(nums)-1])

	if a > b {
		return a
	}

	return b
}
