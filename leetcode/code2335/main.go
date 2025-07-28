package main

import (
	"fmt"
	"sort"
)

func main() {

	// fmt.Println(fillCups([]int{1, 4, 2}))
	fmt.Println(fillCups([]int{1, 1, 1}))
	// fmt.Println(fillCups([]int{0, 0, 0}))
	// fmt.Println(fillCups([]int{5, 0, 0}))
	// fmt.Println(fillCups([]int{0, 0, 2}))
}

func fillCups(amount []int) int {

	var seconds int

	j := 0
	for _, v := range amount {
		if v != 0 {
			amount[j] = v
			j++
		}
	}
	amount = amount[:j]

	for len(amount) != 0 {
		sort.Ints(amount)
		if len(amount) >= 2 {
			n := len(amount)
			amount[n-1]--
			amount[n-2]--
			seconds++
			if amount[n-1] == 0 && amount[n-2] == 0 {
				amount = amount[:n-2]
			} else if amount[n-1] == 0 {
				amount = amount[:n-1]
			} else if amount[n-2] == 0 {
				amount = append(amount[:n-2], amount[n-1])
			}
		} else {
			amount[0]--
			seconds++
			if amount[0] == 0 {
				amount = []int{}
			}
		}
	}

	return seconds
}
