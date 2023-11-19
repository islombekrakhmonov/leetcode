package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(splitNum(999999999))
}

func splitNum(num int) int {
	numStr := strconv.Itoa(num)

	runes := []rune(numStr)

	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	var num1, num2 int
	for i, digit := range runes {
		if i%2 == 0 {
			num1 = num1*10 + int(digit-'0')
		} else {
			num2 = num2*10 + int(digit-'0')
		}
	}

	return num1 + num2
}
