package main

import "fmt"

func main() {
	fmt.Println(minOperations([]int{1,1,1}))
}

func minOperations(nums []int) int {
	countOp := 0

    for i:=0; i<len(nums)-1; i++{
		if nums[i] == nums[i+1] {
			countOp++
			nums[i+1] = nums[i+1] +1 
		} else if nums[i] > nums[i+1] {
			dif := nums[i] - nums[i+1] + 1
			countOp += dif 
			nums[i+1] = nums[i+1] + dif
		}
	}

	return countOp
}