package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(capitalizeTitle("the quiKk brown fo jumps over the lazy dog"))
}

func capitalizeTitle(title string) string {
	var output string

	sliceTitle := strings.Split(title, " ")

	for _, word := range sliceTitle {
		word = CapitalizeFirstLetter(word)
		output += word + " "
	}

	output = strings.TrimRight(output, " ")

	return output
}

func CapitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	} else if len(s) == 1 || len(s) == 2 {
		return strings.ToLower(s)
	}

	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])

	last := s[1:]

	last = strings.ToLower(last)

	output := string(runes[0]) + last

	return output
}
