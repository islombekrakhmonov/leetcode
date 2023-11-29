package main

import (
	"fmt"
)

func main() {
	fmt.Println(largestOddNumber("10133890"))
}

func largestOddNumber(num string) string {
	largestOdd := ""
	for endIndex := len(num); endIndex > 0; endIndex-- {
		subarray := num[:endIndex]
		if isOddNumber(subarray) {
			if len(subarray) > len(largestOdd) {
				largestOdd = subarray
			} else if len(subarray) == len(largestOdd) && subarray > largestOdd {
				largestOdd = subarray
			}
		}
	}
	return largestOdd
}

func isOddNumber(strNum string) bool {
	lastDigit := int(strNum[len(strNum)-1] - '0')
	return lastDigit%2 != 0
}

