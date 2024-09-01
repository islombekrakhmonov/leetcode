package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(pickGifts([]int{25, 64, 9, 4, 100}, 4))
}

func pickGifts(gifts []int, k int) int64 {
	var output int

	for k > 0 {
		max := 0
		maxIndex := 0
		for i, presentNum := range gifts {
			if max < presentNum {
				max = presentNum
				maxIndex = i
			}
		}
		sqRoot := math.Sqrt(float64(max))
		gifts[maxIndex] = int(sqRoot)
		k--
	}

	for _, presentNum := range gifts {
		output += presentNum
	}

	return int64(output)
}
