package main

import "fmt"

func main() {
	fmt.Println(minOperations([]int{3, 1, 5, 4, 2}, 2))
}

func minOperations(nums []int, k int) int {

	occurenceMap := make(map[int]bool)
	for i := len(nums) - 1; i >= 0; i-- {
		if !occurenceMap[nums[i]] && nums[i] <= k {
			occurenceMap[nums[i]] = true
		}

		if len(occurenceMap) == k {
			return len(nums) - i
		}
	}

	return 0
}
