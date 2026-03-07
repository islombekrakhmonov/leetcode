package main

import "fmt"

func main() {
	fmt.Println(residuePrefixes("abc"))
}

func residuePrefixes(s string) int {
	prefix := ""
	count := 0

	for i := 0; i < len(s); i++ {
		prefix += string(s[i])
		if helper(prefix) {
			count++
		}
	}

	return count
}

func helper(s string) bool {
	charCount := make(map[rune]int)

	for _, char := range s {
		charCount[char]++
	}

	if len(charCount) == len(s)%3 {
		return true
	}

	return false
}

//A prefix of s is called a residue if the number of distinct characters in the prefix is equal to len(prefix) % 3.
