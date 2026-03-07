package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(rearrangeCharacters("wvu", "tu"))
}

func rearrangeCharacters(s string, target string) int {
	if len(s) < len(target) {
		return 0
	}

	var count = math.MaxInt
	targetFrequency := make(map[rune]int)
	sFrequency := make(map[rune]int)

	for _, char := range target {
		targetFrequency[char]++
	}

	for _, char := range s {
		if strings.Contains(target, string(char)) {
			sFrequency[char]++
		}
	}

	if len(sFrequency) != len(targetFrequency) {
		return 0
	}

	for char, freq := range sFrequency {
		if targetFrequency[char] == 0 {
			return 0
		}

		if count > freq/targetFrequency[char] {
			count = freq / targetFrequency[char]
		}
	}

	if count == math.MaxInt {
		return 0
	}

	return count
}
