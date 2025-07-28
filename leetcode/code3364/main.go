package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minimumSumSubarray([]int{1, 2, 3, 4, 5}, 2, 4))
}

func minimumSumSubarray(nums []int, l int, r int) int {

	var min = math.MaxInt
	for startIndex := 0; startIndex < len(nums); startIndex++ {
		for endIndex := startIndex + 1; endIndex <= len(nums); endIndex++ {
			subarray := nums[startIndex:endIndex]
			if len(subarray) <= r && len(subarray) >= l {
				var sum int
				for _, num := range subarray {
					sum += num
				}
				if sum > 0 && sum < min {
					min = sum
				}
			}
		}
	}

	if min == math.MaxInt {
		return -1
	}

	return min
}
