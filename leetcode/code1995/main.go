package main

import "fmt"

func main() {
	fmt.Println(countQuadruplets([]int{1, 1, 1, 3, 5}))
}

func countQuadruplets(nums []int) int {

	var output int

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				for l := k + 1; l < len(nums); l++ {
					if nums[i]+nums[j]+nums[k] == nums[l] {
						output++
					}
				}
			}
		}
	}

	return output
}

// Can be written as (a + b) = (d - c)
