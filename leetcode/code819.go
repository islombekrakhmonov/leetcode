package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(mostCommonWord("Bob. hIt, baLl", []string{"bob", "hit"}))
}

func mostCommonWord(paragraph string, banned []string) string {
	lower := strings.ToLower(paragraph)
	removed := removeNonLetters(lower)
	splitted := strings.Split(removed, " ")
	count := make(map[string]int)

	for _, word := range splitted {
		found := false 
		for i:=0; i<len(banned); i++{
			if word == banned[i] {
				found = true
			}
		}
		if !found && word != ""{
			count[word]++
		}
	}

	var max int 
	for _, c := range count {
		if c > max {
			max = c
		}
	}

	for word, c := range count {
		if c == max {
			return word
		}
	}

	return ""
}

func removeNonLetters(input string) string {
	reg := regexp.MustCompile("[^a-zA-Z]")

	result := reg.ReplaceAllString(input, " ")

	return result
}
