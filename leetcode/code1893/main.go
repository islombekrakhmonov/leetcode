package main

import (
	"fmt"
)

func main() {
	// fmt.Println(isCovered([][]int{{1, 2}, {3, 4}, {5, 6}}, 2, 5))
	fmt.Println(isCovered([][]int{{1, 10}, {10, 20}}, 21, 21))
}

func isCovered(ranges [][]int, left int, right int) bool {
	seenMap := make(map[int]bool)

	for i := left; i <= right; i++ {
		seenMap[i] = false
	}

	for _, v := range ranges {
		for i := v[0]; i <= v[1]; i++ {
			seenMap[i] = true
		}
	}

	for _, seen := range seenMap {
		if !seen {
			return false
		}
	}

	return true
}
