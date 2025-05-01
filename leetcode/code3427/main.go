package main

import "fmt"

func main() {
	fmt.Println(subarraySum([]int{2, 3, 1}))
}

func subarraySum(nums []int) int {
	var sum int

	for i := 0; i < len(nums); i++ {
		subArray := nums[max(0, i-nums[i]) : i+1]
		for _, value := range subArray {
			sum += value
		}
	}

	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
