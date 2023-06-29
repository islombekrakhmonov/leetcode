package main

import "fmt"

func main() {
	fmt.Println(mergeAlternately("ab", "pqrs"))
}

func mergeAlternately(word1 string, word2 string) string {
	var output string
    for i:=0; i<len(word1); i++ {
		for l:= i; l<len(word2); l++{
			output += string(word1[i])
			output += string(word2[l])
			break
		}
	}
	if len(word1) > len(word2) {
		output += string(word1[len(word2):])
	}
	if len(word1) < len(word2) {
		output += string(word2[len(word1):])
	}
	return output
}