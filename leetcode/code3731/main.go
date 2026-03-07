package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findMissingElements([]int{5, 7, 11, 13}))
}

func findMissingElements(nums []int) []int {
	min, max := math.MaxInt, math.MinInt
	var frequencyMap = make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if min > nums[i] {
			min = nums[i]
		}

		if max < nums[i] {
			max = nums[i]
		}
		frequencyMap[nums[i]]++
	}

	var output []int

	for i := min + 1; i <= max; i++ {
		if frequencyMap[i] == 0 {
			output = append(output, i)
		}
	}

	return output
}
