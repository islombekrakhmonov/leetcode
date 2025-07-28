package main

import "fmt"

func main() {

	fmt.Println(possibleStringCount("abbcccc"))
}

func possibleStringCount(word string) int {
	frequencyMap := make(map[byte]int)

	var count int

	for i := 0; i < len(word)-1; i++ {
		if word[i] == word[i+1] {
			frequencyMap[word[i]]++
		}
	}

	for _, freq := range frequencyMap {
		count += freq
	}

	return count + 1
}
