package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1,2,3,1},3))
}


func containsNearbyDuplicate(nums []int, k int) bool {
    for i:=0; i<len(nums); i++{
		for l:=i+1;l<len(nums);l++{
			if nums[i] == nums[l] && math.Abs(float64(i) - float64(l)) <=float64(k){
				return true
			}
		}
	}
	return false 
}
