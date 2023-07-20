package main

import "fmt"

func main() {
	fmt.Println(sortedSquares([]int{-4,-1,0,3,10}))
}

func sortedSquares(nums []int) []int {
	var output []int
	for _,v := range nums{
		output = append(output, v*v)
	}

	sorting(output)
	
	
    return output
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