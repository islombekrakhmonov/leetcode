package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(canBeTypedWords("hello world", "id"))
}

func canBeTypedWords(text string, brokenLetters string) int {
	var output int
	textArr := strings.Split(text, " ")

	for _, word := range textArr {
		broken := false
		for _, letter := range brokenLetters {
			if strings.Contains(word, string(letter)) {
				broken = true
			}
		}
		if !broken {
			output++
		}
	}
	return output
}
