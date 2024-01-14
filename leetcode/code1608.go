package main

import (
	"fmt"
)

func main() {
	fmt.Println(specialArray([]int{3,5}))
}

func specialArray(nums []int) int {
	num := 0

	for num <= len(nums) {
		count := 0
		for i := 0; i < len(nums); i++ {
			if nums[i] >= num {
				count++
			}
		}
		if num == count {
			return num
		}
		num++
	}

	return -1
}
