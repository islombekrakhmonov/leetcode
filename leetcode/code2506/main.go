package main

import "fmt"

func main() {
	fmt.Println(similarPairs([]string{"aba", "aabb", "abcd", "bac", "aabc"}))
}

func similarPairs(words []string) int {
	var count int

	wordFreq := make(map[string]map[rune]int)

	for _, word := range words {
		freq := make(map[rune]int)
		for _, char := range word {
			freq[char]++
		}
		wordFreq[word] = freq
	}

	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if isSimilar(wordFreq[words[i]], wordFreq[words[j]]) {
				count++
			}
		}
	}

	return count
}

func isSimilar(freq1, freq2 map[rune]int) bool {
	if len(freq1) != len(freq2) {
		return false
	}

	for char, count := range freq1 {
		if count != freq2[char] {
			return false
		}
	}

	return true
}
