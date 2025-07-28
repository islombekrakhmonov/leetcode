package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(maxFreqSum("successes"))
}

func maxFreqSum(s string) int {
	frequencyMap := make(map[rune]int)

	for _, char := range s {
		frequencyMap[char]++
	}

	var (
		maxCon, maxVow int
	)

	for char, freq := range frequencyMap {
		if isVowel(char) {
			if freq > maxVow {
				maxVow = freq
			}
		} else {
			if freq > maxCon {
				maxCon = freq
			}
		}
	}

	return maxCon + maxVow
}

func isVowel(r rune) bool {
	// Make it lowercase for consistent comparison
	r = unicode.ToLower(r)
	return strings.ContainsRune("aeiou", r)
}
