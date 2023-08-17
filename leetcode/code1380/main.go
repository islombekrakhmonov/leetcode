package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(luckyNumbers([][]int{{3,7,8}, {9,11,13}, {15,16,17}}))
}

func luckyNumbers(matrix [][]int) []int {
	var output []int
	var min = math.MaxInt
	var max = math.MinInt 
    
	for i:=0; i<len(matrix);i++ {
		for j:=0; j < len(matrix[i]);j++{
			if matrix[i][j] < min {
				min = matrix[i][j]
			}
		}

		for k:=0; k<len(matrix); k++{
			if matrix[j][i] > max {
				max = matrix[j][i]
			}
		}

		fmt.Println(max, min)

		if max == min {
			output = append(output, max)
		}

		min = math.MaxInt
		max = math.MinInt
	}
	return output
}