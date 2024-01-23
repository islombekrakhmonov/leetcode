package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(countPairs([]int{-1,1,2,3,1}, 2))
}

func countPairs(nums []int, target int) (output int) {
	sort.Ints(nums)
	left, right := 0,len(nums)-1 

	for left < right {
		if nums[right] + nums[left] < target {
			output += right - left 
			left++
		} else {
			right--
		}
	}
	
	return output
}