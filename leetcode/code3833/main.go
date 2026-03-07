package main

import "fmt"

func main() {
	fmt.Println(dominantIndices([]int{5, 4, 3}))
}

func dominantIndices(nums []int) int {
	var count int

	length := len(nums)
	for startIndex := 0; startIndex < len(nums); startIndex++ {
		if startIndex+1 != length {
			average := average(nums[startIndex+1 : length])
			if nums[startIndex] > average {
				count++
			}
		}
	}

	return count
}

func average(nums []int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}

	return sum / len(nums)
}
