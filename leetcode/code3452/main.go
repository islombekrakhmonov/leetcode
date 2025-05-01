package main

import "fmt"

func main() {
	fmt.Println(sumOfGoodNumbers([]int{1, 3, 2, 1, 5, 4}, 2))
}

func sumOfGoodNumbers(nums []int, k int) int {

	var sum int

	for i := 0; i < len(nums); i++ {
		if i >= k && i+k < len(nums) {
			if nums[i] > nums[i-k] && nums[i] > nums[i+k] {
				sum += nums[i]
			}
		} else if i >= k {
			if nums[i] > nums[i-k] {
				sum += nums[i]
			}
		} else if i+k < len(nums) {
			if nums[i] > nums[i+k] {
				sum += nums[i]
			}

		} else {
			sum += nums[i]
		}
	}

	return sum
}

