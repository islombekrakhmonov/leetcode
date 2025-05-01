package main

import (
	"fmt"
)

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
}

func containsNearbyDuplicate(nums []int, k int) bool {
	indexMap := make(map[int]int)

	for i, num := range nums {
		if lastIndex, found := indexMap[num]; found && i-lastIndex <= k {
			return true
		}
		indexMap[num] = i
	}
	return false
}
