package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findClosest(2, 5, 9))
}

func findClosest(x int, y int, z int) int {
	diffX := absoluteDifference(x, z)
	diffY := absoluteDifference(y, z)

	if diffX < diffY {
		return 1
	} else if diffX > diffY {
		return 2
	} else {
		return 0
	}
}

func absoluteDifference(a, b int) int {
	return int(math.Abs(float64(a - b)))
}
