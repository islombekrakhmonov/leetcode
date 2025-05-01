package main

import "fmt"

func main() {
	fmt.Println(countPartitions([]int{10, 10, 3, 7, 6}))
}

func countPartitions(nums []int) int {

	var (
		n     = len(nums)
		count = 0
	)

	for i := 0; i < n-1; i++ {

		var (
			leftSubArray  = nums[0 : i+1]
			rightSubArray = nums[i+1 : n]

			sumLeft  = 0
			sumRight = 0
		)

		for _, value := range leftSubArray {
			sumLeft += value
		}
		for _, value := range rightSubArray {
			sumRight += value
		}

		if (sumLeft-sumRight)%2 == 0 {
			count++
		}

	}

	return count
}
