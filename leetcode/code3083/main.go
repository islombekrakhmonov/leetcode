package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isSubstringPresent("leetcode"))
}

func isSubstringPresent(s string) bool {

	substings := GenerateSubstrings(s, 2)
	reversed := ReverseString(s)

	for _, substr := range substings {
		if strings.Contains(reversed, substr) {
			return true
		}
	}

	return false
}

func GenerateSubstrings(input string, length int) []string {
	var substrings []string
	strLen := len(input)

	for i := 0; i <= strLen-length; i++ {
		substrings = append(substrings, input[i:i+length])
	}

	return substrings
}

func ReverseString(input string) string {
	runes := []rune(input)
	n := len(runes)

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}

	return string(runes)
}
