package main

import "fmt"

func main() {
	fmt.Println(countPrefixSuffixPairs([]string{"pa", "papa", "ma", "mama"}))
}

func countPrefixSuffixPairs(words []string) int {

	var countPairs int

	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if isPrefixAndSuffix(words[i], words[j]) {
				countPairs++
			}
		}
	}

	return countPairs
}

func isPrefixAndSuffix(str1, str2 string) bool {

	if len(str1) > len(str2) {
		return false
	}

	str1Length := len(str1)

	str2Prefix := str2[0:str1Length]

	str2Suffix := str2[len(str2)-str1Length:]

	return str2Prefix == str1 && str1 == str2Suffix
}
