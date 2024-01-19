package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(deleteGreatestValue([][]int{{1, 2, 4}, {3, 3, 1}}))
}

func deleteGreatestValue(grid [][]int) int {
	var sumOfMax int
	max := make([]int, len(grid[0]))
	for _, v := range grid {
		sort.Ints(v)
		for i := 0; i < len(v); i++ {
			if max[i] < v[i] {
				max[i] = v[i]
			}
		}
	}

	for _, v := range max {
		sumOfMax += v
	}

	return sumOfMax
}
