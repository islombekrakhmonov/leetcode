package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(thirdMax([]int{1,-2147483648,2}))
}

func thirdMax(nums []int) int {
	max :=  math.MinInt
	max2 := math.MinInt
	max3 :=  math.MinInt
	for _,num := range nums {
		if num > max {
			max3 = max2
			max2 = max
			max = num
		} else if max != num && num > max2 {
			max3 = max2
			max2 = num
		} else if max != num && max2 != num && num > max3 {
			max3 = num
		}
	}
	if max3 != math.MinInt {
		return max3
	}

	return max
}