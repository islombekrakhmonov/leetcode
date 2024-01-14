package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(numberGame([]int{1, 2, 3, 4}))
}

func numberGame(nums []int) []int {

	builtInSort(nums)

	for i :=0 ; i<len(nums); i+=2 {
		nums[i], nums[i+1] = nums[i+1], nums[i]
	}
    
	return nums
}

func builtInSort(arr []int)[]int{
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	return arr
}
