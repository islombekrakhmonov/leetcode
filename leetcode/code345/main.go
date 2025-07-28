package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseVowels("IceCreAm"))
}

func reverseVowels(s string) string {
	vowels := []rune{}

	for _, char := range s {
		if isVowel(char) {
			vowels = append(vowels, char)
		}
	}

	reversedVowels := ReverseString(string(vowels))

	output := ""

	for _, char := range s {
		if isVowel(char) {
			output += string(reversedVowels[0])
			reversedVowels = reversedVowels[1:]
		} else {
			output += string(char)
		}
	}

	return output
}

func ReverseString(input string) string {
	runes := []rune(input)
	n := len(runes)

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}

	return string(runes)
}

func isVowel(ch rune) bool {
	vowels := "aeiouAEIOU"
	return strings.ContainsRune(vowels, ch)
}
