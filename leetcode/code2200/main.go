package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findKDistantIndices([]int{3, 4, 9, 1, 3, 9, 5}, 9, 1))
}

func findKDistantIndices(nums []int, key int, k int) []int {
	var output []int

	var keyIndexes []int

	for i := 0; i < len(nums); i++ {
		if nums[i] == key {
			keyIndexes = append(keyIndexes, i)
		}
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(keyIndexes); j++ {
			if int(math.Abs(float64(i-keyIndexes[j]))) <= k {
				output = append(output, i)
				break
			}
		}
	}

	return output
}
