package main

import "fmt"

func main() {
	fmt.Println(repeatedNTimes([]int{1,2,3,3}))
}

func repeatedNTimes(nums []int) int {
	mapOfElements := make(map[int]int)

	for i:=0;i<len(nums);i++{
		mapOfElements[nums[i]]++
	}
	for i, _  := range mapOfElements{
		if mapOfElements[i] == len(nums) / 2{
			return i
		}
	}

	return 0
}