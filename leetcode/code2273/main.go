package main

import "fmt"

func main() {
	fmt.Println(removeAnagrams([]string{"abba", "baba", "bbaa", "cd", "cd"}))
}

func removeAnagrams(words []string) []string {

	for {
		var found bool
		for i := 1; i < len(words); i++ {
			if isAnagram(words[i-1], words[i]) {
				fmt.Println(words[i-1], words[i])
				words = append(words[:i], words[i+1:]...)
				found = true
				break
			}
		}
		if !found {
			break
		}
	}

	return words
}

func isAnagram(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	count := make(map[rune]int)

	for _, char := range word1 {
		count[char]++
	}

	for _, char := range word2 {
		count[char]--
		if count[char] < 0 {
			return false
		}
	}

	return true
}
