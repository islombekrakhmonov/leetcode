package main

import (
	"fmt"
)

func main() {
	fmt.Println(hasTrailingZeros([]int{1, 2, 3, 4, 5}))
}

func hasTrailingZeros(nums []int) bool {
	evenNums := 0

	for _, v := range nums {
		if v%2 == 0 {
			evenNums++
		}
		if evenNums >= 2 {
			return true
		}
	}

	return false
}
