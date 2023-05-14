package main

import "fmt"

func main() {
	nums := []int {7,5,3}
	fmt.Println(smallerNumbersThanCurrent(nums))
}

func smallerNumbersThanCurrent(nums []int) []int {
    var output []int
	var count int
	for i:=0;i<len(nums);i++{
		count = 0
		for p:=0;p<len(nums);p++{
			if nums[i] > nums[p]{
				count ++
			}
		} 
		output = append(output, count)
	}
	return output
}