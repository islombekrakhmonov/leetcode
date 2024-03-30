package main

import "fmt"

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4})) //24,12,8,6
}

func productExceptSelf(nums []int) []int {
	n := len(nums)
	output := make([]int, n)

	leftProduct := 1
	for i := 0; i < n; i++ {
		output[i] = leftProduct
		leftProduct *= nums[i]
	}

	rightProduct := 1
	for i := n - 1; i >= 0; i-- {
		output[i] *= rightProduct
		rightProduct *= nums[i]
	}

	return output
}
