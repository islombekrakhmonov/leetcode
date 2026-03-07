package main

import "fmt"

func main() {

	fmt.Println(maxRepeating("baba", "b"))
}

func maxRepeating(sequence string, word string) int {

	var longest int

	for startIndex := 0; startIndex < len(sequence); startIndex++ {
		for endIndex := startIndex + 1; endIndex <= len(sequence); endIndex++ {
			if endIndex-startIndex < len(word) {
				continue
			}

			substring := sequence[startIndex:endIndex]
			if checkRepeating(substring, word) > longest {
				longest = checkRepeating(substring, word)
			}
		}
	}

	return longest
}

func checkRepeating(substring, word string) int {
	if len(substring) < len(word) {
		return 0
	}

	count := 0
	i := 0

	for i <= len(substring)-len(word) {
		if substring[i:i+len(word)] == word {
			count++
			i += len(word)
		} else {
			break
		}
	}

	return count
}
