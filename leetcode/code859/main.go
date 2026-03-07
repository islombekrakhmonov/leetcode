package main

import "fmt"

func main() {
	fmt.Println(buddyStrings("ab", "ba"))
}

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			s1Runes := []rune(s)
			s1Runes[i], s1Runes[j] = s1Runes[j], s1Runes[i]
			sSwapped := string(s1Runes)
			if sSwapped == goal {
				return true
			}
		}
	}

	return false
}
