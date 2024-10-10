package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(countVowelSubstrings("cuaieuouac"))
}

func countVowelSubstrings(word string) int {

	if len(word) < 5 {
		return 0
	}

	var output int
	var vowels = "aieuo"

	substrings := GenerateSubstrings(word)

	for _, subs := range substrings {
		if isVowelSubstring(subs, vowels) {
			output++
		}
	}

	return output

}

func GenerateSubstrings(input string) []string {
	var substrings []string
	length := len(input)

	for i := 0; i < length; i++ {
		for j := i + 1; j <= length; j++ {
			if len(input[i:j]) > 4 {
				substrings = append(substrings, input[i:j])
			}
		}
	}

	return substrings
}

func isVowelSubstring(subs, vowels string) bool {
	vowelMap := make(map[rune]bool)
	for _, char := range subs {
		if !strings.ContainsRune(vowels, char) {
			return false
		}
		vowelMap[char] = true
	}
	return len(vowelMap) == 5
}
