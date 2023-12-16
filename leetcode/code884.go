package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(uncommonFromSentences("apple apple", "banana"))
}

func uncommonFromSentences(s1 string, s2 string) []string {
	var output []string

	wordCount := make(map[string]int)

	words1 := strings.Fields(s1)
	for _, word := range words1 {
		wordCount[word]++
	}

	words2 := strings.Fields(s2)
	for _, word := range words2 {
		wordCount[word]++
	} 

	fmt.Println(wordCount)

	return output
}
