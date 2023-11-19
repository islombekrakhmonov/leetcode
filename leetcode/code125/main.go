package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}

func isPalindrome(s string) bool {
	clearedStr := strings.Map(func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return -1
		}
		return unicode.ToLower(r)
	}, s)
	strArr := []rune(clearedStr)

	left, right := 0, len(strArr)-1
	for left < right {
		if strArr[left] != strArr[right] {
			return false
		}
		left++
		right--
	}

	return true
}
