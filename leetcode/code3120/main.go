package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(numberOfSpecialChars("aaAbcBC"))
}

func numberOfSpecialChars(word string) int {

	countSpecials := make(map[rune]int)

	upperCaseMap := make(map[string]int)
	lowerCaseMap := make(map[string]int)

	for _, v := range word {
		if unicode.IsLower(v) {
			lowerCaseMap[string(v)]++
		}
		if unicode.IsUpper(v) {
			upperCaseMap[string(v)]++
		}
	}

	for _, value := range word {
		if unicode.IsUpper(value) {
			lowercase := strings.ToLower(string(value))
			if _, exists := lowerCaseMap[lowercase]; exists {
				countSpecials[value]++
				delete(lowerCaseMap, lowercase)
				delete(upperCaseMap, string(value))
			}
		} else {
			upperCase := strings.ToUpper(string(value))
			if _, exists := upperCaseMap[upperCase]; exists {
				countSpecials[value]++
				delete(upperCaseMap, upperCase)
				delete(lowerCaseMap, string(value))
			}
		}
	}

	return len(countSpecials)
}

//You are given a string word. A letter is called special if it appears both in lowercase and uppercase in word.

// Return the number of special letters in word.
