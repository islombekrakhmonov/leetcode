package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countBeautifulPairs([]int{31, 25, 72, 79, 74}))
}

func countBeautifulPairs(nums []int) int {
	var count int

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			num1Str := fmt.Sprintf("%d", nums[i])
			num2Str := fmt.Sprintf("%d", nums[j])

			num1, _ := strconv.Atoi(string(num1Str[0]))
			num2, _ := strconv.Atoi(string(num2Str[len(num2Str)-1]))

			if gcdEuclidean(num1, num2) == 1 {
				count++
			}
		}
	}

	return count
}

func gcdEuclidean(a, b int) int {
	if b == 0 {
		return a
	}
	return gcdEuclidean(b, a%b)
}
