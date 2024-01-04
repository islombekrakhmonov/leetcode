package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(sortVowels("lEetcOde"))
}

func sortVowels(s string) string {
	vowels := make([]rune, 0)
	output := make([]rune, len(s))

	for i,v := range s {
		if isVowel(v) {
			vowels = append(vowels, v)
			output[i] = -1 
		} else {
			output[i] = v
		}
	}

	sort.Slice(vowels, func(i, j int) bool {
		return vowels[i] < vowels[j]
	})

	vowelIndex := 0

	for i, ch := range output {
		if ch == -1 {
			output[i] = vowels[vowelIndex]
			vowelIndex++
		}
	}
	
	return string(output)
}

func isVowel(ch rune) bool {
	vowels := "aeiouAEIOU"
	return strings.ContainsRune(vowels, ch)
}
