package main

import (
	"fmt"
)

func main() {
	fmt.Println(targetIndices([]int{4,2,5,2,1}, 2))
}

func targetIndices(nums []int, target int) []int {
	var ourput []int

	sorted := bubbleSort(nums)

	for i,v := range sorted{
		if v == target{
			ourput = append(ourput, i)
		}
	}
    
	return ourput
}

func bubbleSort(arr []int) []int{
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}