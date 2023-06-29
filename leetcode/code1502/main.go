package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(canMakeArithmeticProgression([]int{1,100}))
}

func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	firstDifference := math.Abs(float64(arr[1] - arr[0]))
	for i:=1; i<len(arr)-1; i++{
		if firstDifference != math.Abs(float64(arr[i+1] - arr[i])){
			return false
		}
	}
	return true
}