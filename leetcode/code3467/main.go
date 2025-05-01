package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(transformArray([]int{4, 3, 2, 1}))
}

func transformArray(nums []int) []int {

	for i, v := range nums {
		if v%2 != 0 {
			nums[i] = 1
		} else {
			nums[i] = 0
		}
	}

	sort.Ints(nums)

	return nums
}
