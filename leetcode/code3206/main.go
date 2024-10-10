package main

import "fmt"

func main() {
	fmt.Println(numberOfAlternatingGroups([]int{0, 1, 0, 0, 1}))
}

func numberOfAlternatingGroups(colors []int) int {

	var count int
	l := len(colors)
	for i := 0; i < l; i++ {
		previousIndex := (i - 1 + l) % l
		nextIndex := (i + 1) % l

		if colors[previousIndex] == colors[nextIndex] && colors[previousIndex] != colors[i] {
			count++
		}
	}

	return count
}
