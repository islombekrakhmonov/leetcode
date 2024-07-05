package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(toGoatLatin("I speak Goat Latin"))
}

func toGoatLatin(sentence string) string {
	var output string

	sliceSentence := strings.Split(sentence, " ")

	for i, word := range sliceSentence {
		word = convertWord(word, i+1)
		fmt.Println(word)
		output += word + " "
	}

	output = strings.TrimRight(output, " ")

	return output
}

func convertWord(s string, index int) string {
	for i, letter := range s {
		if i == 0 && !isVowel(letter) {
			s = s[1:]
			s += string(letter)
		}
	}

	s += "ma"

	for index != 0 {
		s += "a"
		index--
	}

	return s
}

func isVowel(ch rune) bool {
	vowels := "aeiou"
	return strings.ContainsRune(vowels, unicode.ToLower(ch))
}
