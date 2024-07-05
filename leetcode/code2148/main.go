package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(countElements([]int{11, 7, 2, 15}))
}

func countElements(nums []int) int {
	var count int
	var max = math.MinInt
	var min = math.MaxInt

	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
		if nums[i] < min {
			min = nums[i]
		}
	}

	for _, v := range nums {
		if v == max || v == min {
			count++
		}
	}

	return count
}
