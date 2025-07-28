package main

import "fmt"

func main() {
	fmt.Println(constructTransformedArray([]int{-1, 4, -1}))
}

func constructTransformedArray(nums []int) []int {
	l := len(nums)

	var result = make([]int, l)

	for i := 0; i < l; i++ {
		if nums[i] > 0 {
			result[i] = nums[(i+nums[i])%l]
		} else if nums[i] < 0 {
			landedIndex := (i + nums[i]) % l
			if landedIndex < 0 {
				landedIndex += l
			}
			result[i] = nums[landedIndex]
		} else if nums[i] == 0 {
			result[i] = 0
		}
	}

	return result
}

//If nums[i] > 0: Start at index i and move nums[i] steps to the right in the circular array. Set result[i] to the value of the index where you land.
// If nums[i] < 0: Start at index i and move abs(nums[i]) steps to the left in the circular array. Set result[i] to the value of the index where you land.
// If nums[i] == 0: Set result[i] to nums[i].
