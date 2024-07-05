package main

import "fmt"

func main() {
	fmt.Println(commonChars([]string{"cool", "lock", "cook"}))
}

func commonChars(words []string) []string {
	var output []string

	minFreq := make(map[rune]int)
	for _, char := range words[0] {
		minFreq[char]++
	}

	for _, word := range words[1:] {
		wordFreq := make(map[rune]int)

		for _, char := range word {
			wordFreq[char]++
		}
		for char := range minFreq {
			if wordFreq[char] < minFreq[char] {
				minFreq[char] = wordFreq[char]
			}
		}

		for char := range minFreq {
			if wordFreq[char] == 0 {
				delete(minFreq, char)
			}
		}

	}

	for char, count := range minFreq {
		for i := 0; i < count; i++ {
			output = append(output, string(char))
		}
	}

	return output
}
