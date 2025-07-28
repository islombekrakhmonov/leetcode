package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(minDeletion("aaaaaaaaabbbbbbbbbbbccccccccjjjjjjj", 2))
}

func minDeletion(s string, k int) int {
	var (
		frequencyMap  = make(map[rune]int)
		countArray    = []int{}
		countDeletion int
	)

	for _, char := range s {
		frequencyMap[char]++
	}

	if len(frequencyMap)-k <= 0 {
		return 0
	}

	for _, v := range frequencyMap {
		countArray = append(countArray, v)
	}

	sort.Ints(countArray)

	for len(countArray) != k {
		countDeletion += countArray[0]
		countArray = countArray[1:]
	}

	return countDeletion
}
