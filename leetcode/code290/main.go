package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(wordPattern("abba", "dog dog dog dog"))
}

func wordPattern(pattern string, s string) bool {
	sSlice := strings.Split(s, " ")

	if len(pattern) != len(sSlice) {
		return false
	}

	mappingFromPatter := make(map[rune]string)
	mappingFromS := make(map[string]rune)

	for i, letter := range pattern {
		if _, exists := mappingFromPatter[letter]; exists {
			if mappingFromPatter[letter] != sSlice[i] {
				return false
			}
		} else if _, exists := mappingFromS[sSlice[i]]; exists {
			if mappingFromS[sSlice[i]] != letter {
				return false
			}
		} else {
			mappingFromPatter[letter] = sSlice[i]
			mappingFromS[sSlice[i]] = letter
		}
	}

	return true
}
