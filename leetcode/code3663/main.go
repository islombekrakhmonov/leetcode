package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(getLeastFrequentDigit(11))
}

func getLeastFrequentDigit(n int) int {
	var digits []int
	var frequency = make(map[int]int)
	minFreq := 100
	for n > 0 {
		num := n % 10
		frequency[num]++
		n /= 10
	}

	for _, freq := range frequency {
		if freq < minFreq {
			minFreq = freq
		}
	}

	for num, freq := range frequency {
		if freq == minFreq {
			digits = append(digits, num)
		}
	}

	sort.Ints(digits)

	return digits[0]
}
