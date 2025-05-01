package main

import (
	"fmt"
)

func main() {
	fmt.Println(canBeIncreasing([]int{1, 2, 5, 4, 5}))

	price := 130.52

	totalPax := 4

	price *= float64(totalPax)
	fmt.Println("price:", price)
	// fmt.Println("nums:", nums)

	// nums = append(nums[:1], nums[1+1:]...)
	// fmt.Println("nums", nums)
}

func canBeIncreasing(nums []int) bool {

	if isIncreasing(nums) {
		return true
	}

	for i := 0; i < len(nums); i++ {
		array := make([]int, 0, len(nums)-1)
		array = append(array, nums[:i]...)
		array = append(array, nums[i+1:]...)
		if isIncreasing(array) {
			return true
		}
	}

	return false
}

func isIncreasing(nums []int) bool {

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] >= nums[i+1] {
			return false
		}
	}

	return true
}
