package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(missingMultiple([]int{1, 4, 3, 2, 5}, 3))
}

func missingMultiple(nums []int, k int) int {

	sort.Ints(nums)
	var maxInt = nums[len(nums)-1]
	var possibles []int

	for i := k; i <= maxInt+k; i += k {
		possibles = append(possibles, i)
	}

	for _, v := range possibles {
		found := false
		for _, v2 := range nums {
			if v == v2 {
				found = true
				break
			}
		}
		if !found {
			return v
		}
	}

	return 0
}
