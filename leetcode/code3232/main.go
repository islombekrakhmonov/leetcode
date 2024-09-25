package main

import "fmt"

func main() {
	fmt.Println(canAliceWin([]int{
		1, 2, 3, 4, 5, 14,
	}))
}

func canAliceWin(nums []int) bool {
	var single, double int

	for _, num := range nums {
		if num >= -9 && num <= 9 {
			single += num
		} else {
			double += num
		}
	}

	if single == double {
		return false
	}

	return true
}
