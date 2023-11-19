package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(detectCapitalUse("USA"))
}

func detectCapitalUse(word string) bool {
	var capital, lower = 0,0

	for _, v := range word {
		ifCapital := unicode.IsUpper(v)
		if ifCapital {
			capital++
		} else {
			lower++
		}
	}

	if lower == len(word) || capital == len(word) {
		return true
	} else if unicode.IsUpper(rune(word[0])) && lower == len(word) - 1 {
		return true
	}
	

    return false 
}
