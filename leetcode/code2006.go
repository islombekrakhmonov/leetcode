package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(countKDifference([]int{1,2,2,1}, 1))
}

func countKDifference(nums []int, k int) int {
	var output int 

    for i:=0; i<len(nums); i++{
		for j:=i+1; j<len(nums); j++{
			if int(math.Abs(float64(nums[i] - nums[j]))) == k {
				output++
			}
		}
	}
	return output
}