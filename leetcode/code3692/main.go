package main

import (
	"fmt"
)

func main() {
	fmt.Println(majorityFrequencyGroup("pfpfgi"))
}

func majorityFrequencyGroup(s string) string {
	var output string
	maxGroupSize := 0
	maxGroupFreq := 0

	frequencyMap := make(map[string]int)
	groupFrequencyMap := make(map[int]int)

	for _, v := range s {
		frequencyMap[string(v)]++
	}

	for _, v := range frequencyMap {
		groupFrequencyMap[v]++
	}

	for k, v := range groupFrequencyMap {
		if (v > maxGroupSize) || (v == maxGroupSize && k > maxGroupFreq) {
			maxGroupSize = v
			maxGroupFreq = k
		}
	}

	for k, v := range frequencyMap {
		if v == maxGroupFreq {
			output += k
		}
	}

	return output
}
