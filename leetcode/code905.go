package main

import "fmt"

func main() {
	fmt.Println(sortArrayByParity([]int{3,2,4,5}))	
}

func sortArrayByParity(nums []int) []int {
	l := len(nums)
    for i:=0; i<l; i++{
		for j:=0; j<l-i-1; j++{
			if nums[j] % 2 != 0 {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}