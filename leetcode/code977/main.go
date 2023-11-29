package main

import "fmt"

func main() {
	fmt.Println(sortedSquares([]int{-4,-1,0,3,10}))
}

func sortedSquares(nums []int) []int {
	for i,v := range nums{
		nums[i] = v*v
	}

	sorting(nums)
	
	
    return nums
}

func sorting(nums []int) []int{
	l := len(nums)

	for i:=0; i<l;i++{
		for j:=0; j<l-i-1; j++{
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1],nums[j]
			}
		}
	}
	return nums
}