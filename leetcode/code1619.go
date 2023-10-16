package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(trimMean([]int{6,0,7,0,7,5,7,8,3,4,0,7,8,1,6,8,1,1,2,4,8,1,9,5,4,3,8,5,10,8,6,6,1,0,6,10,8,2,3,4}))
}

func trimMean(arr []int) float64 {

	sort.Ints(arr)
	n := len(arr)
	removeCount := n / 20

	sum := 0
	for i := removeCount; i < n-removeCount; i++ {
		sum += arr[i]
	}

	fmt.Println( float64(n-2*removeCount))

	mean := float64(sum) / float64(n-2*removeCount)
	roundedResult := math.Round(mean*1e5) / 1e5


	return roundedResult
}