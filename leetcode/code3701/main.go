package main

import "fmt"

func main() {
	fmt.Println(alternatingSum([]int{3, 1, 2, 10, 0}))
}

func alternatingSum(nums []int) int {
	var sum int

	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			sum += nums[i]
		} else {
			sum -= nums[i]
		}
	}

	return sum
}
