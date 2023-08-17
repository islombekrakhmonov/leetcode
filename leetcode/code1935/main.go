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
	textArr := strings.Split(text," ")

    Loop:
	for _, word := range textArr {
		for _, letter := range brokenLetters {
			if strings.ContainsRune(word, letter) {
				continue Loop
			}
		}
		output++
	}
	return output
}