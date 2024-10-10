package main

import (
	"fmt"
)

func main() {
	fmt.Println(countGoodRectangles([][]int{{5, 8}, {3, 9}, {5, 12}, {16, 5}}))
}

func countGoodRectangles(rectangles [][]int) int {

	var squares []int
	var output int

	maxLen := 0

	for _, sides := range rectangles {
		min := min(sides[0], sides[1])
		squares = append(squares, min)
		if min > maxLen {
			maxLen = min
		}
	}

	for _, v := range squares {
		if v == maxLen {
			output++
		}
	}

	return output
}

func min(num1, num2 int) int {
	if num1 > num2 {
		return num2
	}

	return num1
}
