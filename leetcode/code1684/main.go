package main

import (
	"fmt"
	"strings"
	// "strings"
)

func main() {
	fmt.Println(countConsistentStrings("ab", []string{"ad","bd","aaab","baa","badab"}))
}

func countConsistentStrings(allowed string, words []string) int {
	var count int 
	var notContains int
	for _, v := range words {
		notContains = 0
		for i,l := range v {
			if !strings.Contains(allowed, string(l)){
				notContains ++
				fmt.Println(string(l), "at index", i)
			}
		}
		if notContains == 0 {
			count++
		}
	}
	return count
}