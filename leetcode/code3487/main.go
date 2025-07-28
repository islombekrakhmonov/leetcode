package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxSum([]int{-1}))
}

func maxSum(nums []int) int {

	frequencyMap := make(map[int]bool)
	var sum int
	var max = math.MinInt

	for _, v := range nums {
		if v >= 0 {
			frequencyMap[v] = true
		}
		if max < v {
			max = v
		}
	}

	if len(frequencyMap) == 0 {
		return max
	}

	for char := range frequencyMap {
		sum += char
	}

	return sum
}
