package main

import "fmt"

func main() {
	fmt.Println(countLargestGroup(13))
}

func countLargestGroup(n int) int {
	count := make(map[int]int)
	for i := 1; i <= n; i++ {
		sum := 0
		for j := i; j > 0; j /= 10 {
			sum += j % 10
		}
		count[sum]++
	}
	maxCount := 0
	for _, v := range count {
		if v > maxCount {
			maxCount = v
		}
	}

	output := 0
	for _, v := range count {
		if v == maxCount {
			output++
		}
	}

	return output
}
