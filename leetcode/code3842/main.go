package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(toggleLightBulbs([]int{10, 30, 20, 10}))
}

func toggleLightBulbs(bulbs []int) []int {
	countMap := make(map[int]int)

	for _, num := range bulbs {
		countMap[num]++
	}

	var output []int
	for bulb, count := range countMap {
		if count%2 == 1 {
			output = append(output, bulb)
		}
	}

	sort.Ints(output)

	return output
}
