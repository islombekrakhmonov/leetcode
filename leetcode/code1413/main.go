package main

import (
	"fmt"
)

func main() {
	fmt.Println(minStartValue([]int{-3, 2, -3, 4, 2}))
}

func minStartValue(nums []int) int {

	minPrefixSum := 0
	prefixSum := 0

	// Calculate the prefix sum and track the minimum
	for _, num := range nums {
		prefixSum += num
		if prefixSum < minPrefixSum {
			minPrefixSum = prefixSum
		}
	}

	return 1 - minPrefixSum
}
