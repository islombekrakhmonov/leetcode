package main

import "fmt"

func main() {
	fmt.Println(largestLocal([][]int{{9, 9, 8, 1}, {5, 6, 2, 6}, {8, 2, 6, 4}, {6, 2, 2, 2}}))
}

func largestLocal(grid [][]int) [][]int {
	n := len(grid)
	output := make([][]int, len(grid)-2)
	for i := 0; i < len(grid)-2; i++ {
		output[i] = make([]int, n-2)
		for j := 1; j < len(grid)-2; j++ {
			output[i][j] = findMax(grid, i, j)
		}
	}

	return output
}

func findMax(grid [][]int, x, y int) int {
	maxElement := grid[x][y]
	for i := x; i < x+3; i++ {
		for j := y; j < y+3; j++ {
			if grid[i][j] > maxElement {
				maxElement = grid[i][j]
			}
		}

	}
	return maxElement
}

/*
Algorithm
Create an empty matrix maxLocal of size (N−2)⋅(N−2)(N - 2) \cdot (N - 2)(N−2)⋅(N−2), this will store the maximum values of all possible 3 x 3 matrices.
Define the findMax function, which takes the grid and the coordinates (x, y) as parameters. This function finds the maximum value in the 3 x 3 section of the grid, where (x, y) is the top-left corner.
Iterate over the 3 x 3 matrix starting with (x, y) as top-left cell.
Find and return the maximum value as maxElement.
Iterate over the grid rows 0 to N - 2 and columns 0 to N - 2, and for each cell (i, j):
Use findMax(grid, i, j) to find the maximum local element and store it in the matrix maxLocal at position (i, j).
Return maxLocal.
*/
