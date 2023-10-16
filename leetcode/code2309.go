package main

import (
	"fmt"
	"strings"
	"time"
	"unicode"
)
func main() {
	s := "arRAzFif"
	
	// Measure the runtime of GreatestLetter
	start := time.Now()
	result := GreatestLetter(s)
	elapsed := time.Since(start)
	fmt.Printf("GreatestLetter result: %s\n", result)
	fmt.Printf("GreatestLetter runtime: %s\n", elapsed)

	// Measure the runtime of GreatestLetter1
	start = time.Now()
	result = GreatestLetter1(s)
	elapsed = time.Since(start)
	fmt.Printf("GreatestLetter1 result: %s\n", result)
	fmt.Printf("GreatestLetter1 runtime: %s\n", elapsed)
}

func GreatestLetter(s string) string {
	var characters []byte
	for i:=0; i<len(s); i++{
		count := 0
		for j := i+1; j<len(s); j++{
			unicodeCodePoint := int(s[i])
			if unicode.IsLower(rune(s[i])) {
				if int(s[j]) == unicodeCodePoint -32 {
					count ++
				}
			} else if unicode.IsUpper(rune(s[i])) {
				if int(s[j]) == unicodeCodePoint +32 {
					count ++
				}
			}
		}
		if count >= 1 {
			if unicode.IsLower(rune(s[i])) {
				characters = append(characters,byte(rune(s[i]-32)) )
			} else {
				characters = append(characters, s[i])
			}
		}
	}
	if len(characters) > 0{
		max := characters[0]
		for _,v := range characters {
			if max < v {
				max = v
			}
		}
		return (string(rune(max)))
	}
	
    
	return ""
}

func GreatestLetter1(s string) string {
	// Create a map to store lowercase letters that occur in s
	lowercase := make(map[rune]bool)

	// Initialize the result to an empty string
	result := ""

	// Iterate through the string s
	for _, char := range s {
		// Check if the character is an uppercase letter
		if unicode.IsUpper(char) {
			// Check if the lowercase version of the character exists in the map
			if lowercase[unicode.ToLower(char)] {
				// Update the result if the current character is greater
				if result == "" || char > rune(result[0]) {
					result = string(char)
				}
			}
		} else if unicode.IsLower(char) {
			// If the character is lowercase, add it to the map
			lowercase[char] = true
		}
	}

	// Convert the result to uppercase and return it
	if result != "" {
		return strings.ToUpper(result)
	}

	// If no matching character was found, return an empty string
	return ""
}