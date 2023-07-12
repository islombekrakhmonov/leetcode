package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(leftRightDifference([]int{10,4,8,3}))
}

func leftRightDifference(nums []int) []int {
	var answer []int 
	
	
	for i:=0; i<len(nums); i++{
		rightSum := 0
		for j:=i+1; j<len(nums); j++{
			rightSum += nums[j]
		}
		leftSum := 0
		for k:=i-1; k>=0; k--{
			leftSum += nums[k]
		}
		
		absolute := math.Abs(float64(leftSum-rightSum))

		answer = append(answer,int(absolute))
	}
	
	return answer
}