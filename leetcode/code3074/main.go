package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumBoxes([]int{5, 5, 5}, []int{2, 4, 2, 7}))
}

func minimumBoxes(apple []int, capacity []int) int {

	var countApples int
	var countBoxes int

	bags := 0
	output := 0

	for _, v := range apple {
		countApples += v
	}

	sort.Sort(sort.Reverse(sort.IntSlice(capacity)))

	for _, v := range capacity {
		bags += v
		countBoxes++
		if bags >= countApples {
			output = countBoxes
			break
		}
	}

	return output
}
