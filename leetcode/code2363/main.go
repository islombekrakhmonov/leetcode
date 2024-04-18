package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(mergeSimilarItems([][]int{{1, 1}, {4, 5}, {3, 8}}, [][]int{{3, 1}, {1, 5}}))
}

func mergeSimilarItems(items1, items2 [][]int) [][]int {
	merged := make([][]int, 0)

	// Merge items from items1
	for _, item := range items1 {
		merged = mergeItem(merged, item)
	}

	// Merge items from items2
	for _, item := range items2 {
		merged = mergeItem(merged, item)
	}

	// Sort the merged items by value
	sort.Slice(merged, func(i, j int) bool {
		return merged[i][0] < merged[j][0]
	})

	return merged
}

// Helper function to merge an item into the merged slice
func mergeItem(merged [][]int, item []int) [][]int {
	for i, existing := range merged {
		if existing[0] == item[0] {
			merged[i][1] += item[1] // Add weight to existing item
			return merged
		}
	}
	// If item not found, append it to the merged slice
	return append(merged, item)
}
