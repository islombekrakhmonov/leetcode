package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumAverage([]int{7, 8, 3, 4, 15, 13, 4, 1}))
}

func minimumAverage(nums []int) float64 {
	var averages []float64

	i := len(nums) / 2

	sort.Ints(nums)

	for i > 0 {
		averages = append(averages, (float64(nums[0]+nums[len(nums)-1]) / 2))
		nums = nums[:len(nums)-1]
		nums = nums[1:]
		i--
	}

	sort.Float64s(averages)

	return averages[0]

}
