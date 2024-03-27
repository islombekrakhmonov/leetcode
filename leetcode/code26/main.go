package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}

func removeDuplicates(nums []int) int {
	var left = 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[left] = nums[i]
			left++
		}
	}
	return left
}
