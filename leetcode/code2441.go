package main

import "fmt"

func main() {
	fmt.Println(findMaxK([]int{-1,2,-3,3}))
}

func findMaxK(nums []int) int {
	output := -1 
	for i:=0; i<len(nums); i++ {
		for j:=0; j<len(nums); j++ {
			if nums[i] == -nums[j] {
				if nums[i] > 0  && nums[i] > output {
					output = nums[i]
				} else if nums[j] > 0 && nums[j] > output {
					output = nums[j]
				}
			}
		}
	}
	return output
}