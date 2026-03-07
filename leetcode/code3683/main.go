package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(earliestTime([][]int{{100, 100}, {100, 100}, {100, 100}}))
}

func earliestTime(tasks [][]int) int {
	var times []int

	for _, task := range tasks {
		times = append(times, task[0]+task[1])
	}

	sort.Ints(times)

	return times[0]
}
