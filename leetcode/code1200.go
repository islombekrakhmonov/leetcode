package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(minimumAbsDifference([]int{3,8,-10,23,19,-4,-14,27}))
}

func minimumAbsDifference(arr []int) [][]int {
	minDifference := math.MaxInt
	sort.Ints(arr)
	var output [][]int

	fmt.Println("sorted", arr)

	for i := 0; i < len(arr)-1; i++ {
        diff := int(math.Abs(float64(arr[i+1]-arr[i])))
        if diff < minDifference {
			minDifference = diff
            output = [][]int{{arr[i], arr[i+1]}}
			fmt.Println(minDifference)
			fmt.Println("diff", output)
        } else if diff == minDifference {
			fmt.Println("before", output)
            output = append(output, []int{arr[i], arr[i+1]})
			fmt.Println("output", output)
        }
    }

	return output
}