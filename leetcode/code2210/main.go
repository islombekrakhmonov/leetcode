package main

import "fmt"

func main() {
	fmt.Println(countHillValley([]int{2, 4, 1, 1, 6, 5}))
}

func countHillValley(nums []int) int {

	var output int

	for i := 1; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] != nums[j] {
				if nums[i] > nums[i-1] && nums[i] > nums[j] {
					output++
				} else if nums[i] < nums[i-1] && nums[i] < nums[j] {
					output++
				}
				break
			}
		}
	}

	return output
}
