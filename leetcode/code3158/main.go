package main

import "fmt"

func main() {
	fmt.Println(duplicateNumbersXOR([]int{1, 2, 2, 1}))
}

func duplicateNumbersXOR(nums []int) int {
	var xor int

	countMap := make(map[int]int)

	for _, v := range nums {
		countMap[v]++
	}

	for i, v := range countMap {
		if v == 2 {
			xor ^= i
		}
	}

	return xor
}
