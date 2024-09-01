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

	countSpecials := make(map[byte]int)

	for i := 0; i < len(word); i++ {
		for j := i + 1; j < len(word); j++ {
			if unicode.IsUpper(rune(word[j])) && string(word[i]) == strings.ToLower(string(word[j])) {
				countSpecials[word[i]]++
			} else if unicode.IsLower(rune(word[j])) && string(word[i]) == strings.ToUpper(string(word[j])) {
				countSpecials[word[i]]++
			}
		}
	}

	return len(countSpecials)
}
