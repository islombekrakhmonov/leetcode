package main

import "fmt"

func main() {
	fmt.Println(maxFrequencyElements([]int{10, 12, 11, 9, 6, 19, 11}))
}

func maxFrequencyElements(nums []int) int {
	var output int
	frequency := make(map[int]int)
	for _, v := range nums {
		frequency[v]++
	}

	maxFreq := 0
	for _, v := range frequency {
		if v > maxFreq {
			maxFreq = v
		}
	}

	for _, v := range frequency {
		if maxFreq == v {
			output += v
		}
	}

	return output
}
