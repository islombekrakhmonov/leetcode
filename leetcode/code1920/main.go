package main

import "fmt"

func main() {
	fmt.Println(buildArray([]int{0,2,1,5,3,4}))
}

func buildArray(nums []int) []int {
	var output []int
    
	for i:=0; i<len(nums); i++{
		output = append(output, nums[nums[i]])
	}

	return output
}