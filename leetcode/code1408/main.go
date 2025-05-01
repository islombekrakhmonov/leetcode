package main

import "fmt"

func main() {
	fmt.Println(stringMatching([]string{"leetcoder", "leetcode", "od", "hamlet", "am"}))
}

func stringMatching(words []string) []string {
	var output []string
	seen := make(map[string]bool)

	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			isSubstring, substr := containsSubstring(words[i], words[j])
			if !seen[substr] {
				if isSubstring {
					output = append(output, substr)
					seen[substr] = true
				}
			}
		}
	}

	return output
}

func containsSubstring(word1, word2 string) (bool, string) {
	var word, substr string

	if len(word1) >= len(word2) {
		word = word1
		substr = word2
	} else {
		word = word2
		substr = word1
	}

	l := len(substr)
	for i := 0; i <= len(word)-l; i++ {
		if word[i:i+l] == substr {
			return true, substr
		}
	}

	return false, ""
}
