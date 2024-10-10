package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(arrayRankTransform([]int{100, 100, 100}))
}

func arrayRankTransform(arr []int) []int {
	sortedArr := make([]int, len(arr))
	copy(sortedArr, arr)

	sort.Ints(sortedArr)

	rankMap := make(map[int]int)
	rank := 1

	for _, val := range sortedArr {
		if _, exists := rankMap[val]; !exists {
			rankMap[val] = rank
			rank++
		}
	}

	ranks := make([]int, len(arr))
	for i, val := range arr {
		ranks[i] = rankMap[val]
	}

	return ranks
}
