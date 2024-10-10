package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxDistance([]int{9, 9, 9, 18, 9, 9, 9, 9, 9, 18}))
}

func maxDistance(colors []int) int {
	var max int

	for i, color := range colors {
		for j := i + 1; j < len(colors); j++ {
			if color != colors[j] {
				if j-i > max {
					max = j - i
				}
			}
		}
	}

	return max
}
