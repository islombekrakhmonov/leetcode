package main

import "fmt"

func main() {
	fmt.Println(maximumLengthSubstring("bcbbbcba"))
}

func maximumLengthSubstring(s string) int {
	var output int
	for startIndex := 0; startIndex < len(s); startIndex++ {
		for endIndex := startIndex + 1; endIndex <= len(s); endIndex++ {
			substring := s[startIndex:endIndex]
			length := countLength(substring)
			if length > output {
				output = length
			}
		}
	}

	return output
}

func countLength(substring string) int {
	frequencyMap := make(map[byte]int)
	for i := 0; i < len(substring); i++ {
		frequencyMap[substring[i]]++
		if frequencyMap[substring[i]] > 2 {
			return 0
		}
	}

	return len(substring)
}
