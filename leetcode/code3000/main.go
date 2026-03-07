package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(areaOfMaxDiagonal([][]int{{9, 3}, {8, 6}}))
}

func areaOfMaxDiagonal(dimensions [][]int) int {

	maxDiagonal := 0.0
	maxArea := 0

	for i := 0; i < len(dimensions); i++ {
		diagonal := math.Sqrt(float64(dimensions[i][0]*dimensions[i][0] + dimensions[i][1]*dimensions[i][1]))
		if diagonal > maxDiagonal {
			maxDiagonal = diagonal
			maxArea = dimensions[i][0] * dimensions[i][1]
		} else if diagonal == maxDiagonal {
			area := dimensions[i][0] * dimensions[i][1]
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}
