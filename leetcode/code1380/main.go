package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(luckyNumbers([][]int{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}}))
}

func luckyNumbers(matrix [][]int) []int {
	var output []int

	for row := 0; row < len(matrix); row++ {
		rowMin := math.MaxInt
		colIndex := 0

		for col := 0; col < len(matrix[row]); col++ {
			if matrix[row][col] < rowMin {
				rowMin = matrix[row][col]
				colIndex = col
			}
		}

		isMaxInCol := true
		for i := 0; i < len(matrix); i++ {
			if matrix[i][colIndex] > rowMin {
				isMaxInCol = false
				break
			}
		}

		if isMaxInCol {
			output = append(output, rowMin)
		}

	}

	return output
}
