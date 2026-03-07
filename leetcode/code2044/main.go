package main

import "fmt"

func main() {
	fmt.Println(countMaxOrSubsets([]int{3, 2, 1, 5}))
}

func countMaxOrSubsets(nums []int) int {

	var (
		maxOr int
	)

	subsets := subsets(nums)
	orMap := make(map[int]int)

	for _, array := range subsets {
		var bitwiseOr int
		for _, num := range array {
			bitwiseOr |= num
		}
		orMap[bitwiseOr]++
		if bitwiseOr > maxOr {
			maxOr = bitwiseOr
		}
	}

	return orMap[maxOr]
}

func subsets(nums []int) [][]int {
	var result [][]int
	n := len(nums)
	total := 1 << n // 2^n possible subsets

	for mask := 0; mask < total; mask++ {
		var subset []int
		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				subset = append(subset, nums[i])
			}
		}
		result = append(result, subset)
	}
	return result
}
