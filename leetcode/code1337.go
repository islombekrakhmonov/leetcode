package main

import (
	"fmt"
	"sort"
)


func main() {
	mat := [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1},
	}
	fmt.Println(kWeakestRows(mat, 3))
}

type RowInfo struct {
    Index       int
    SoldierCount int
}


func kWeakestRows(mat [][]int, k int) []int {
	rowInfos := make([]RowInfo, len(mat))
	
	for i, row := range mat {
        soldierCount := binarySearch(row)
		rowInfos[i] = RowInfo{Index: i, SoldierCount: soldierCount}
	}

	sort.Slice(rowInfos, func(i, j int) bool {
        if rowInfos[i].SoldierCount == rowInfos[j].SoldierCount {
            return rowInfos[i].Index < rowInfos[j].Index
        }
        return rowInfos[i].SoldierCount < rowInfos[j].SoldierCount
    })

	output := make([]int, k)
    for i := 0; i < k; i++ {
        output[i] = rowInfos[i].Index
    }

	return output
}

func binarySearch(arr []int) int {
	left, right := 0, len(arr)
	for left < right {
		mid := left + (right-left)/2
		if arr[mid] < 1 {
			right = mid
		} else {
            left = mid + 1
        }
	}
	return left
}
