package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1,3,5,6}, 7))
}


func searchInsert(nums []int, target int) int {
	for i:=0; i<len(nums); i++{
		if nums[i] == target {
			return i
		} else if i < len(nums)-1 && nums[i] < target && nums[i+1] > target {
			return i+1
		}
	}

	if nums[len(nums)-1] < target {
		return len(nums)
	}

	return 0
}