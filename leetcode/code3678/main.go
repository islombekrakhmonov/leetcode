package main

import "fmt"

func main() {
	fmt.Println(smallestAbsent([]int{-34}))
}

func smallestAbsent(nums []int) int {
	var (
		sum      int
		smallest int
	)

	for _, num := range nums {
		sum += num
	}

	smallest = int(float64(sum)/float64(len(nums)) + 1)

	if smallest < 1 {
		smallest = 1
	}

	for {
		if !contains(nums, smallest) {
			break
		}
		smallest++
	}

	return smallest
}

func contains(slice []int, element int) bool {
	for _, el := range slice {
		if el == element {
			return true
		}
	}
	return false
}
