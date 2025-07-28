package main

import "fmt"

func main() {
	fmt.Println(countSubarrays([]int{-1, -4, -1, 4}))
}

func countSubarrays(nums []int) int {

	var count int
	for i := 0; i < len(nums)-2; i++ {
		if 2*(nums[i]+nums[i+2]) == nums[i+1] {
			count++
		}
	}

	return count
}
