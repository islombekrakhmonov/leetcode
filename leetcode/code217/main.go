package main

import "fmt"

func main() {
	fmt.Println(containsDuplicate([]int{1,2,3,4,5,2}))
}

func containsDuplicate(nums []int) bool {
    for i:=0; i<len(nums); i++{
		for l:=i+1; l<len(nums); l++{
			if nums[i] == nums[l] {
				return true
			} 
		}
	}
	return false 
}