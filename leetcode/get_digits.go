package main

import "fmt"

func main() {
	fmt.Println(getDigits(345343))
}

func getDigits(number int) []int {
    if number == 0 {
        return []int{0}
    }

	var digits []int
	for number > 0 {
		digit := number % 10
		digits = append(digits, digit)
		number /= 10
	}

	// Reverse the slice
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	return digits
}