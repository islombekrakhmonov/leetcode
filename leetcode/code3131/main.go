package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(addedInteger([]int{2, 6, 4}, []int{9, 7, 5}))
}

func addedInteger(nums1 []int, nums2 []int) int {
	var x int

	sort.Ints(nums1)
	sort.Ints(nums2)

	if len(nums1) > 0 && len(nums2) > 0 {
		x = nums2[0] - nums1[0]
	}

	return x
}
