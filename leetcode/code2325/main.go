package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(decodeMessage("the quick brown fox jumps over the lazy dog", "vkbs bs t suepuv"))
}

func decodeMessage(key string, message string) string {

	var output string

	uniqueChars := make(map[rune]bool)

	keyTrimmed  := strings.ReplaceAll(key, " ", "")

	result := ""
	for _, char := range keyTrimmed {
		if !uniqueChars[char] {
			uniqueChars[char] = true 
			result += string(char)
		}
	}
	
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	alphabetSlice := strings.Split(alphabet, "")

	alphabet1 := make(map[string]string)

	for i,v := range result {
		alphabet1[string(v)] = alphabetSlice[i]
	}
	for _,v := range message {
		if string(v) == " "{
			output += " "
		}
		output += alphabet1[string(v)]
	}

	return output
}