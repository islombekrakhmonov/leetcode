package main

import "fmt"

func main() {
	permutations := permute([]int{1, 2, 3})
	fmt.Println(permutations)
}
func permute(nums []int) [][]int {
	var result [][]int
	var backtrack func(int)

	backtrack = func(first int) {
		// if all positions fixed, add a copy to results
		if first == len(nums) {
			perm := make([]int, len(nums))
			copy(perm, nums)
			result = append(result, perm)
			return
		}

		for i := first; i < len(nums); i++ {
			// swap current element with the first
			nums[first], nums[i] = nums[i], nums[first]
			// recursively fix the rest
			backtrack(first + 1)
			// backtrack (undo swap)
			nums[first], nums[i] = nums[i], nums[first]
			fmt.Println(nums)
		}
	}

	backtrack(0)
	return result
}
