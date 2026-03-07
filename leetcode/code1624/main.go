package main

import "fmt"

func main() {
	fmt.Println(maxLengthBetweenEqualCharacters("abca"))
}

func maxLengthBetweenEqualCharacters(s string) int {
	firstSeen := make(map[byte]int)
	maxLen := -1

	for i := 0; i < len(s); i++ {
		if idx, ok := firstSeen[s[i]]; ok {
			length := i - idx - 1
			if length > maxLen {
				maxLen = length
			}
		} else {
			firstSeen[s[i]] = i
		}
	}

	return maxLen
}
