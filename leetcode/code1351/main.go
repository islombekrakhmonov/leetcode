package main

import "fmt"

func main() {
	fmt.Println(countNegatives([][]int{{4,3,2,-1},{3,2,1,-1},{1,1,-1,-2},{-1,-1,-2,-3}}))
}

func countNegatives(grid [][]int) int {
    totalCount := 0
    for _, row := range grid {
        totalCount += countNegativesInRow(row)
		fmt.Println(totalCount)
    }
    return totalCount
}

func countNegativesInRow(row []int) int {
    left, right := 0, len(row)
    for left < right {
        mid := left + (right-left)/2
        if row[mid] < 0 {
            right = mid
        } else {
            left = mid + 1
        }
	}
    return len(row) - left
}
