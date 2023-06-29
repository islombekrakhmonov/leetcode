package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}

func isPalindrome(s string) bool {
	nonAlphanumericRegex := regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
	removed := nonAlphanumericRegex.ReplaceAllString(s, "")
	trimmedS := strings.ReplaceAll(removed, " ", "")
	toLower := strings.ToLower(trimmedS)
	var backward string

	for i := len(toLower); i > 0; i-- {
		backward += string(toLower[i-1])
	}
	fmt.Println(backward)
	return toLower == backward
}
