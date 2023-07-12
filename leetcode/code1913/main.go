package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxProductDifference([]int{10,10,10,10}))
}

func maxProductDifference(nums []int) int {
	max1 := math.MinInt64
	max2 := math.MinInt64
	min1 := math.MaxInt64
	min2 := math.MaxInt64


	for _,v := range nums{
		if max1 <= v {
			max2 = max1
			max1 = v
		} else if v > max2 && max1 != v{
			max2 = v
		}
		if min1 >= v {
			min2 = min1
			min1 = v
		} else if min2 > v && min1 != v {
			min2 = v
		}
	}
	fmt.Println(max1,max2)

    return (max1*max2)-(min1*min2)
}