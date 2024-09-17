package main

import (
	"fmt"
)

func main() {
	fmt.Println(countConsistentStrings("ab", []string{"ad", "bd", "aaab", "baa", "badab"}))
}

func countConsistentStrings(allowed string, words []string) int {

	var count int
	charMap := make(map[rune]bool)

	for _, char := range allowed {
		charMap[char] = true
	}

	for _, str := range words {
		var charCount int
		for _, char := range str {
			if _, hasChar := charMap[char]; hasChar {
				charCount++
			}
		}
		if charCount == len(str) {
			count++
		}
	}

	return count
}
