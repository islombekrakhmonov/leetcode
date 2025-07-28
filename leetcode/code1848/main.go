package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(getMinDistance([]int{1, 2, 3, 4, 5}, 5, 3))
}

func getMinDistance(nums []int, target int, start int) int {
	min := math.MaxInt

	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			abs := int(math.Abs(float64(i - start)))
			if abs < min {
				min = abs
			}
		}
	}

	return min
}
