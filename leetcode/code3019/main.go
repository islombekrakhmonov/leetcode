package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(countKeyChanges("AaAaAaaA"))
}

func countKeyChanges(s string) int {
	var count int

	for i := 1; i < len(s); i++ {
		if unicode.ToLower(rune(s[i-1])) != unicode.ToLower(rune(s[i])) {
			count++
		}
	}

	return count
}

