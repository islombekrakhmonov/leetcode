package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minMovesToSeat([]int{3, 1, 5}, []int{2, 7, 4}))
}

func minMovesToSeat(seats []int, students []int) int {

	var output int
	sort.Ints(seats)
	sort.Ints(students)

	for i, v := range students {
		dif := v - seats[i]

		if dif < 0 {
			dif = -dif
		}

		output += dif
	}

	return output
}
