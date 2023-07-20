package main

import "fmt"

func main() {
	fmt.Println(sumOfUnique([]int{1,2,3,2}))
}

func sumOfUnique(nums []int) int {
    uniqueNumbers := make(map[int]int)
	var output int
	for _,v:= range nums{
		uniqueNumbers[v]++
	}

	for i:=0;i<len(nums);i++{
		if uniqueNumbers[nums[i]] == 1{
			output += nums[i]
		}
	}
	return output
}