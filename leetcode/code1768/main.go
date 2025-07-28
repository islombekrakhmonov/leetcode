package main

import "fmt"

func main() {
	fmt.Println(mergeAlternately("ab", "pqrs"))
}

func mergeAlternately(word1 string, word2 string) string {
	var output string

	longerStr := ""
	if len(word1) > len(word2) {
		longerStr = word1
	} else {
		longerStr = word2
	}

	for i := 0; i < len(longerStr); i++ {
		if i < len(word1) {
			output += string(word1[i])
		}
		if i < len(word2) {
			output += string(word2[i])
		}
	}

	return output
}
