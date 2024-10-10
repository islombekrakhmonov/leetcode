package main

import "fmt"

func main() {
	fmt.Println(returnToBoundaryCount([]int{2, 3, -5}))
}

func returnToBoundaryCount(nums []int) int {

	var (
		count int
		steps int
	)

	for _, num := range nums {
		steps += num

		if steps == 0 {
			count++
		}
	}

	return count
}
