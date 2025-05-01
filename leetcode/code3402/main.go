package main

import "fmt"

func main() {
	fmt.Println(minimumOperations([][]int{{3, 2}, {1, 3}, {3, 4}, {0, 1}}))
}

func minimumOperations(grid [][]int) int {
	var minOperation int

	for i := 0; i < len(grid)-1; i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] >= grid[i+1][j] {
				operation := grid[i][j] - grid[i+1][j] + 1
				grid[i+1][j] += operation
				minOperation += operation
			}
		}
	}

	return minOperation
}
