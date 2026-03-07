package main

import "fmt"

func main() {
	fmt.Println(reversePrefix("abcd", 2))
}

func reversePrefix(s string, k int) string {

	if len(s) >= k {
		return string(reverseSlice([]rune(s[:k]))) + s[k:]
	}

	return string(reverseSlice([]rune(s)))
}

func reverseSlice(slice []rune) []rune {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}
