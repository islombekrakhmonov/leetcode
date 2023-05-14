package main

import (
	"fmt"
)

func main() {
	a := []int{2,3,5}
	fmt.Println(listNum(a,7))
	
}

func listNum(nums []int, target int)int{
	var final []int 
	var x int 
	for i,v := range nums {
		if v == target{
			x = i
		} else {
			nums = append(nums, target)
			for i=0;i<len(nums);i++{
				if target > nums[i]{
					final = append(final, nums[i])
				} else if target <nums[i] && target > nums[i-1]{
					final = append(final, nums[i])
				} else if target < nums[i]{
					final = append(final, nums[i])
				}
			}
		}
		for i,v := range final{
			if target == v {
				x = i
				return x
			}
		}
	}
	return x
}