package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(countPrefixes([]string{"a","b","c","ab","bc","abc"}, "abc"))
}

func countPrefixes(words []string, s string) int {
	var output int
	for i := 0; i<len(words); i++{
		if strings.HasPrefix(s, words[i]) {
			output ++
		}
	}
    
	return output
}