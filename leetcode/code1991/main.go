package main

import "fmt"

func main() {
	fmt.Println(findMiddleIndex([]int{2, 3, -1, 8, 4}))
}

func findMiddleIndex(nums []int) int {
	totalSum := 0
	for _, v := range nums {
		totalSum += v
	}

	leftSum := 0

	for i := 0; i < len(nums); i++ {
		if i != 0 {
			leftSum += nums[i-1]
		}
		if totalSum-leftSum-nums[i] == leftSum {
			return i
		}
	}

	return -1

}

//Find the total sum first and then iterate again to check whether a index is middleindex or not.
