package main

import "fmt"

func main() {
	fmt.Println(largestLocal([][]int{{9, 9, 8, 1}, {5, 6, 2, 6}, {8, 2, 6, 4}, {6, 2, 2, 2}}))
}

func largestLocal(grid [][]int) [][]int {
	output := make([][]int, len(grid)-2)
	for i := 0; i < len(grid); i++ {
		row := make([]int, len(grid)-2)
		for j := 1; j < len(grid)-1; j++ {
			curMax := grid[i][j]
			fmt.Println(curMax)
			fmt.Println(row)
		}
	}

	return output
}
