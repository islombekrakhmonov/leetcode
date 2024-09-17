package main

import "fmt"

func main() {
	fmt.Println(isMonotonic([]int{1, 2, 2, 3}))
}

func isMonotonic(nums []int) bool {

	isMonotoneInc, isMonotoneDec := true, true

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			isMonotoneDec = false
		}

		if nums[i] < nums[i+1] {
			isMonotoneInc = false
		}

		if !isMonotoneDec && !isMonotoneInc {
			return false
		}
	}

	return true
}
