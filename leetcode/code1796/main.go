package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(secondHighest("sjhtz8344"))
}

func secondHighest(s string) int {
	max := math.MinInt
	secondMax := math.MinInt

	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			num, _ := strconv.Atoi(string(s[i]))
			fmt.Println(max, secondMax, num)
			if num > max {
				secondMax = max
				max = num
			} else if secondMax < num && num < max {
				secondMax = num
			}
		}
	}

	if secondMax > math.MinInt {
		return secondMax
	}

	return -1
}
