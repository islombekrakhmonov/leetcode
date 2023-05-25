package main

import "fmt"

func main() {
	fmt.Println(runningSum([]int{1,2,3,4}))
}

func runningSum(nums []int) []int {
    var output []int
	var number int
	for i:= 0; i<len(nums); i++{
		number = 0
		for j:= 0; j<i+1; j++{
			number += nums[j]
		}
		output = append(output, number)
	}
	return output
}