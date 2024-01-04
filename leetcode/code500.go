package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(findWords([]string{"Hello","Alaska","Dad","Peace"}))
}

func findWords(words []string) []string {
	var output []string
	first := "qwertyuiop"
	second := "asdfghjkl"
	third := "zxcvbnm"

	for i:=0; i<len(words); i++ {
		toLower := strings.ToLower(words[i])
		found := true
		for j:=0; j<len(toLower); j++{
			if !strings.Contains(first, string(toLower[j])) {
				found = false
			}
		}
		if found {
			output = append(output, words[i])
		}
		found = true
		for j:=0; j<len(toLower); j++{
			if !strings.Contains(second, string(toLower[j])) {
				found = false
			}
		}
		if found {
			output = append(output, words[i])
		}

		found = true
		for j:=0; j<len(toLower); j++{
			if !strings.Contains(third, string(toLower[j])) {
				found = false
			}
		}
		if found {
			output = append(output, words[i])
		}
	}
	
	return output
}