package main

import (
	"fmt"
)

func main() {
	fmt.Println(topKFrequent([]int{4, 1, -1, 2, -1, 2, 3}, 2))
}

func topKFrequent(nums []int, k int) []int {
	mapped := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		mapped[nums[i]]++
	}

	bucket := make([][]int, len(nums)+1)
	
	for key, val := range mapped {
		bucket[val] = append(bucket[val], key)
	}

	ans := make([]int, 0, k)

	for i := len(bucket) - 1; i >= 0; i-- {
		for _, val := range bucket[i] {
			if k > 0 {
				ans = append(ans, val)
				k--
			} else if k == 0 {
				break
			}
		}
	}

	return ans
}
