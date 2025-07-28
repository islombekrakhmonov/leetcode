package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(mostFrequentEven([]int{0, 1, 2, 2, 4, 4, 1}))
}

func mostFrequentEven(nums []int) int {

	var (
		frequencyMap = make(map[int]int)
		maxFrequency = 0
		min          = math.MaxInt
	)

	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			frequencyMap[nums[i]]++
			frequency := frequencyMap[nums[i]]
			if frequency > maxFrequency {
				maxFrequency = frequency
			}
		}
	}

	for num, freq := range frequencyMap {
		if freq == maxFrequency {
			if num < min {
				min = num
			}
		}
	}
	if min == math.MaxInt {
		return -1
	}

	return min
}
