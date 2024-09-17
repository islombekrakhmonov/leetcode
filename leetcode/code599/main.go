package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findRestaurant([]string{"happy", "sad", "good"}, []string{"sad", "happy", "good"}))
}

func findRestaurant(list1 []string, list2 []string) []string {
	var output []string

	hashMap := make(map[string]int)
	minSum := math.MaxInt

	for i := 0; i < len(list1); i++ {
		for j := 0; j < len(list2); j++ {
			if list1[i] == list2[j] {
				sum := i + j
				if _, exists := hashMap[list1[i]]; !exists {
					hashMap[list1[i]] = sum
					if exists && hashMap[list1[i]] > i+j {
						hashMap[list1[i]] = sum
					}
					if sum < minSum {
						minSum = sum
					}
				}
			}
		}
	}

	for i, v := range hashMap {
		if v == minSum {
			output = append(output, i)
		}
	}

	return output
}
