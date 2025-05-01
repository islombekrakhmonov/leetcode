package main

import (
	"fmt"
	"sort"
)

func main() {
	// [8,7],[9,9],[7,4],[9,7]
	fmt.Println(maxWidthOfVerticalArea([][]int{{8, 7}, {9, 9}, {7, 4}, {9, 7}}))
}

func maxWidthOfVerticalArea(points [][]int) int {
	array := []int{}

	for i := range points {
		array = append(array, points[i][0])
	}

	sort.Slice(array, func(a, b int) bool {
		return array[a] > array[b]
	})

	maxDifference := 0
	for i := 0; i < len(array)-1; i++ {
		difference := array[i] - array[i+1]
		if difference > maxDifference {
			maxDifference = difference
		}
	}

	return maxDifference
}

// Get every zero index points[i][0] elements in to array
