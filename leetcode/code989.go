package main

import (
	"fmt"
)

func main() {
	fmt.Println(addToArrayForm([]int{1, 2, 6, 3, 0, 7, 1, 7, 1, 9, 7, 5, 6, 6, 4, 4, 0, 0, 6, 3}, 516))
}

func addToArrayForm(num []int, k int) []int {
	ans := make([]int, 0)
	for i := len(num) - 1; k > 0 || i >= 0; {
		sum := k % 10
		k /= 10
		if i >= 0 {
			sum += num[i]
			i--
		}
		ans = append(ans, sum%10)
		k += sum / 10
	}
	reverseSlice(ans)

	return ans
}

func reverseSlice(slice []int) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}
