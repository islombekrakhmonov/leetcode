package main

import "fmt"

func main() {
	fmt.Println(sumDivisibleByK([]int{1, 2, 2, 3, 3, 3, 3, 40}, 2))
}

func sumDivisibleByK(nums []int, k int) int {
	frequencyMap := make(map[int]int)

	for _, num := range nums {
		frequencyMap[num]++
	}

	var output int

	for num, freq := range frequencyMap {
		if freq%k == 0 {
			output += num * freq
		}
	}

	return output
}
