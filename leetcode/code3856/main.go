package main

import "fmt"

func main() {
	fmt.Println(trimTrailingVowels("aeiou"))
}

func trimTrailingVowels(s string) string {

	var count int
	for i := len(s) - 1; i >= 0; i-- {
		if !isVowel(s[i]) {
			return s[:i+1]
		} else {
			count++
		}
	}

	if count == len(s) {
		return ""
	}

	return s
}

func isVowel(char byte) bool {
	vowelChar := []byte{'a', 'e', 'i', 'o', 'u'}
	for _, v := range vowelChar {
		if char == v {
			return true
		}
	}
	return false
}
