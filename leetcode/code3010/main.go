package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(minimumCost([]int{5, 4, 3}))
}

func minimumCost(nums []int) int {
	var sum int

	sum += nums[0]
	nums = nums[1:]
	sort.Ints(nums)

	sum += nums[0] + nums[1]

	return sum
}
