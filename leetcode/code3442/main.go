package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxDifference("abcabcab"))
}

func maxDifference(s string) int {

	frequncyMap := make(map[rune]int)
	var maxOdd int
	var minEven = math.MaxInt

	for _, char := range s {
		frequncyMap[char]++
	}

	for _, freq := range frequncyMap {
		if freq%2 != 0 {
			if maxOdd < freq {
				maxOdd = freq
			}
		} else {
			if minEven > freq {
				minEven = freq
			}
		}
	}

	return maxOdd - minEven
}
