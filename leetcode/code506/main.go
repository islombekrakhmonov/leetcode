package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(findRelativeRanks([]int{10, 3, 8, 9, 4}))

	s := []int{0, 1, 2, 3, 4, 5}
	copy(s[2:], s[0:])
	fmt.Println(s)
}

func findRelativeRanks(score []int) []string {
	ranks := make([]string, len(score))

	indices := make(map[int]int)
	for i, v := range score {
		indices[v] = i
	}

	sort.SliceStable(score, func(a, b int) bool {
		return score[a] > score[b]
	})

	for i := 0; i < len(score); i++ {
		index := indices[score[i]]
		if i == 0 {
			ranks[index] = "Gold Medal"
		} else if i == 1 {
			ranks[index] = "Silver Medal"
		} else if i == 2 {
			ranks[index] = "Bronze Medal"
		} else {
			ranks[index] = strconv.Itoa(i + 1)
		}
	}

	return ranks
}
