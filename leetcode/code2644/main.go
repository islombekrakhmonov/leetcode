package main

import "fmt"

func main() {
	fmt.Println(maxDivScore([]int{69, 3, 92, 14, 67, 90, 31, 40, 54, 63, 99, 88, 28, 100, 5, 72, 89, 60, 90, 71}, []int{97, 16, 7, 60, 6, 57, 73, 84, 17, 8, 77, 60, 7, 74, 74, 24, 52, 43, 94, 48, 9, 99}))
}

func maxDivScore(nums []int, divisors []int) int {
	maxScore := -1
	result := divisors[0]

	for _, divisor := range divisors {
		score := 0
		for _, num := range nums {
			if num%divisor == 0 {
				score++
			}
		}

		if score > maxScore || (score == maxScore && divisor < result) {
			maxScore = score
			result = divisor
		}
	}

	return result
}
