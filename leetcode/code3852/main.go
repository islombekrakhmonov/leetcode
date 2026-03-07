package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minDistinctFreqPair([]int{1, 1, 2, 2, 3, 4}))
}

func minDistinctFreqPair(nums []int) []int {
	freq := make(map[int]int)

	x := math.MaxInt
	xFreq := 0

	for _, num := range nums {
		if x > num {
			x = num
		}
		freq[num]++
		if num == x {
			xFreq = freq[num]
		}
	}

	var y, yFreq int
	for num, freq := range freq {
		if num != x && freq != xFreq {
			if y == 0 {
				y = num
				yFreq = freq
			} else if num < y {
				y = num
				yFreq = freq
			} else if num == y && freq < yFreq {
				y = num
				yFreq = freq
			}
		}
	}

	if y != 0 {
		return []int{x, y}
	}

	return []int{-1, -1}
}
