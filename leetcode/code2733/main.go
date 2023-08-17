package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findNonMinOrMax([]int{2,4,25}))
}

func findNonMinOrMax(nums []int) int {
	min := math.MaxInt
	max := math.MinInt
	for _,v := range nums{
		if max < v {
			max = v
		}
		if min > v {
			min=v
		}
	}

	for _,v := range nums{
		if v > min && max > v {
			return v
		}
	}
    
	return -1
}