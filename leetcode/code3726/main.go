package main

import "fmt"

func main() {
	fmt.Println(removeZeros(9180))
}

func removeZeros(x int64) int64 {
	var output int64
	multiplier := 1

	for x > 0 {
		digit := x % 10
		if digit != 0 {
			output += int64(digit) * int64(multiplier)
			multiplier *= 10
		}
		x /= 10
	}
	return output
}
