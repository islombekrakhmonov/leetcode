package main

import "fmt"

func main() {
	fmt.Println(buddyStrings("ab", "ba"))
}

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	// if s == goal {
	// 	seen := make(map[rune]bool)
	// 	for i,v := range
	// }

	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			char1 := s[i]
			char2 := s[j]

			s[i] = char2
			s[j] = char1 
			
		}
	}

	return false
}
