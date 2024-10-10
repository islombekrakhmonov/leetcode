package main

import "fmt"

func main() {
	fmt.Println(findLucky([]int{2, 2, 3, 4}))
}

func findLucky(arr []int) int {
	frequencyMap := make(map[int]int)

	for _, v := range arr {
		frequencyMap[v]++
	}

	max := -1

	for i, v := range frequencyMap {
		if i == v && max < i {
			max = i
		}
	}

	return max
}
