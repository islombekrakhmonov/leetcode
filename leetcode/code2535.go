package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int {1,15,6,3}
	fmt.Println(differenceOfSum(nums))
	fmt.Println(16/10)
}

func differenceOfSum(nums []int) int {
    var elementSum int
	var digitSum int
	var difference int

	for _,v := range nums{
		elementSum += v
		for v>0{
			digitSum += v%10
			v /=10
		}
	}
	difference = int(math.Abs(float64(elementSum)-float64(digitSum)))
	
	return difference
}