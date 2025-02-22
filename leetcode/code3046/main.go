package main

import "fmt"

func main() {
	fmt.Println(isPossibleToSplit([]int{1, 1, 2, 2, 3, 4}))
}

func isPossibleToSplit(nums []int) bool {

	countMap := make(map[int]int)

	for _, v := range nums {
		countMap[v]++
	}

	for _, count := range countMap {
		if count > 2 {
			return false
		}
	}

	return true
}
