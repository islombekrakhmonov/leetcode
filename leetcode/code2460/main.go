package main

import "fmt"

func main() {
	fmt.Println(applyOperations([]int{1, 2, 2, 1, 1, 0}))
}

func applyOperations(nums []int) []int {

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			nums[i] = nums[i] * 2
			nums[i+1] = 0
		}
	}

	var output []int

	var countZeros int

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			output = append(output, nums[i])
		} else {
			countZeros++
		}
	}

	for countZeros != 0 {
		output = append(output, 0)
		countZeros--
	}

	return output
}
