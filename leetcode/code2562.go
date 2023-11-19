package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(findTheArrayConcVal([]int{7, 52, 2, 4}))
}

func findTheArrayConcVal(nums []int) int64 {
	var concatenation int
	for len(nums) > 0 {
		if len(nums) > 1 {
			firstElement := strconv.Itoa(nums[0])
			secondElement := strconv.Itoa(nums[len(nums)-1])
			sum, _ := strconv.Atoi(firstElement + secondElement)
			concatenation += sum
			nums = nums[1:len(nums)-1]
		} else {
			concatenation += nums[0]
			nums = []int{}
		}
	}

	return int64(concatenation)
}
