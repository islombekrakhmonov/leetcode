package main

import "fmt"

func main() {
	fmt.Println(maxProduct([]int{3,4,5,2}))
}

func maxProduct(nums []int) int {
	l := len(nums)
	for i:=0;i<l;i++{
		for j:=0;j<l-i-1;j++{
			if nums[j+1] > nums[j]{
				nums[j], nums[j+1] = nums[j+1], nums[j]
			} 
		}
	}

	output := (nums[0]-1)*(nums[1]-1)
    
	return output
}