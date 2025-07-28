package main

import "fmt"

func main() {
	fmt.Println(maxOperations([]int{3, 2, 1, 4, 5}))
}

func maxOperations(nums []int) int {
	var count int
	var target int

	if len(nums) > 1 {
		target = nums[0] + nums[1]
		nums = nums[2:]
		count++
	}

	for {
		if len(nums) < 2 {
			return count
		}
		if target != nums[0]+nums[1] {
			return count
		} else {
			nums = nums[2:]
			count++
		}
	}
}
