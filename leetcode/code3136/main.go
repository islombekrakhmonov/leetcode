package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(isValid("234Adas"))
}

func isValid(word string) bool {

	if len(word) < 3 {
		return false
	}

	var validStringRegex = regexp.MustCompile(`^[a-zA-Z0-9]+$`)

	if !validStringRegex.MatchString(word) {
		return false
	}

	lowerInput := strings.ToLower(word)

	hasVowel := false
	hasConsonant := false

	vowels := "aeiou"

	for _, char := range lowerInput {
		if strings.ContainsRune(vowels, char) {
			hasVowel = true
		} else if char >= 'a' && char <= 'z' {
			hasConsonant = true
		}

		if hasVowel && hasConsonant {
			return true
		}
	}

	return false
}
