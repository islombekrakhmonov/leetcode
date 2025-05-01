package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxAdjacentDistance([]int{1, 3, 4, 9, 2}))
}

func maxAdjacentDistance(nums []int) int {

	maxAdj := 0

	for i := 0; i < len(nums); i++ {
		var abs float64

		if i == len(nums)-1 {
			abs = math.Abs(float64(nums[0]) - float64(nums[i]))
		} else {
			abs = math.Abs(float64(nums[i]) - float64(nums[i+1]))
		}
		if abs > float64(maxAdj) {
			maxAdj = int(abs)
		}
	}

	return maxAdj
}
