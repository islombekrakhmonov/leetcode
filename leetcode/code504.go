package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(convertToBase7(-7))
}

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}

	isNegative := false
	if num < 0 {
		isNegative = true
		num = -num
	}

	var result string
	for num > 0 {
		digit := num % 7
		result = strconv.Itoa(digit) + result
		num /= 7
	}

	if isNegative {
		result = "-" + result
	}

	return result
}