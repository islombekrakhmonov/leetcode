package main

import "fmt"

func main() {
	fmt.Println(countPairs([]int{-1,1,2,3,1}, 2))
}

func countPairs(nums []int, target int) (output int) {
    n := len(nums) 

	for i :=0; i<n; i++{
		for j:=i+1; j<n; j++{
			if nums[i] + nums[j] < target{
				output++
			}
		}
	}
	return output
}