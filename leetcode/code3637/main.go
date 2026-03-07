package main

import "fmt"

func main() {
	fmt.Println(isTrionic([]int{5, 9, 1, 7}))
}

func isTrionic(nums []int) bool {

	for i := 1; i < len(nums)-2; i++ {
		if isIncreasing(nums[:i+1]) {
			for j := i + 1; j < len(nums)-1; j++ {
				if isDecreasing(nums[i : j+1]) {
					if isIncreasing(nums[j:]) {
						return true
					}
				}
			}
		}
	}

	return false
}

func isIncreasing(nums []int) bool {
	if len(nums) < 2 {
		return false
	}

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] >= nums[i+1] {
			return false
		}
	}
	return true
}

func isDecreasing(nums []int) bool {
	if len(nums) < 2 {
		return false
	}

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] <= nums[i+1] {
			return false
		}
	}
	return true
}
