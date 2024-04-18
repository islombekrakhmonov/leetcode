package main

import "fmt"

func main() {

	fmt.Println(findErrorNums([]int{1, 2, 2, 4}))
}

func findErrorNums(nums []int) []int {
	n := len(nums)
	seen := make([]bool, n+1)
	dup := 0

	for _, num := range nums {
		if seen[num] {
			dup = num
		}
		seen[num] = true
	}

	for i := 1; i <= n; i++ {
		if !seen[i] {
			return []int{dup, i}
		}
	}

	return []int{dup, 0}
}
