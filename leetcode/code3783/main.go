package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(mirrorDistance(25))
}

func mirrorDistance(n int) int {
	nStr := strconv.Itoa(n)

	nStr = string(reverseSlice([]rune(nStr)))
	reversed, _ := strconv.Atoi(nStr)

	return int(math.Abs(float64(n - reversed)))
}

func reverseSlice(slice []rune) []rune {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}
