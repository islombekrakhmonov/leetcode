package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minMoves([]int{1, 2, 3}))
	agentAdultNetPrice := math.Ceil(49.305*10) / 10

	fmt.Println(agentAdultNetPrice)
}

func minMoves(nums []int) int {
	maxInt := math.MinInt

	for _, num := range nums {
		if maxInt < num {
			maxInt = num
		}
	}

	moves := 0

	for _, num := range nums {
		moves += maxInt - num
	}

	return moves
}
