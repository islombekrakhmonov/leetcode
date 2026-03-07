package main

import (
	"fmt"
)

func main() {
	fmt.Println(sumAndMultiply(10203004))
}

func sumAndMultiply(n int) int64 {
	var sum int
	var reversed int

	for n > 0 {
		digit := n % 10
		sum += digit
		if digit != 0 {
			reversed = reversed*10 + digit
		}
		n /= 10
	}

	result := 0
	for reversed > 0 {
		result = result*10 + (reversed % 10)
		reversed /= 10
	}

	return int64(result) * int64(sum)
}
