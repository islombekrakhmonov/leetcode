package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findPermutationDifference("abc", "bac"))
}

func findPermutationDifference(s string, t string) int {

	var difference int

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(t); j++ {
			if s[i] == t[j] {
				absolute := math.Abs(float64(i) - float64(j))
				difference += int(absolute)
			}
		}
	}

	return difference
}
