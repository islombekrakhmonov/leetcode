package main

import "fmt"

func main() {
	fmt.Println(minimumOperations([]int{4, 5, 6, 4, 4}))
}

func minimumOperations(nums []int) int {
	operations := 0

	for {
		disctinct := true
		frequencyMap := make(map[int]int)
		for i := 0; i < len(nums); i++ {
			frequencyMap[nums[i]]++
			if frequencyMap[nums[i]] > 1 {
				if len(nums) > 3 {
					nums = nums[3:]
					disctinct = false
				} else {
					nums = []int{}
				}
				operations++
				break
			}
		}
		if disctinct {
			break
		}
	}

	return operations
}
