package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(lastStoneWeight([]int{2, 2}))
}

func lastStoneWeight(stones []int) int {

	for {

		if len(stones) == 0 {
			return 0
		}

		if len(stones) == 1 {
			return stones[0]
		}

		sort.Slice(stones, func(i, j int) bool {
			return stones[i] > stones[j]
		})

		fmt.Println(stones)
		if stones[0] != stones[1] {
			stones = append(stones, stones[0]-stones[1])
		}

		stones = stones[2:]
	}
}
