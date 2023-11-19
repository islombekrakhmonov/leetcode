package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findTheDistanceValue([]int{1,4,2,3}, []int{-4,-3,6,10,20,30}, 3))
}

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
    var output int
	for i :=0; i<len(arr1); i++ {
		violates := 0
		for j:=0; j<len(arr2); j++ {
			difference := int(math.Abs(float64(arr1[i] - arr2[j])))
			if difference <= d {
				violates++
			}
		}
		if violates ==0 {
			output++
		}
	}
	return output
}