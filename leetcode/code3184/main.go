package main

import "fmt"

func main() {
	fmt.Println(countCompleteDayPairs([]int{12, 12, 30, 24, 24}))
}

func countCompleteDayPairs(hours []int) int {
	var completeDays int

	for i := 0; i < len(hours)-1; i++ {
		for j := i + 1; j < len(hours); j++ {
			if (hours[i]+hours[j])%24 == 0 {
				completeDays++
			}
		}
	}

	return completeDays

}
