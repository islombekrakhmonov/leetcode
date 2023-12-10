package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(findWordsContaining([]string{"leet","code"}, 'e'))
}

func findWordsContaining(words []string, x byte) []int {
	var output []int

	for i:=0; i<len(words); i++{
		if strings.Contains(words[i], string(x)) {
			output = append(output, i)
		}
 	}
    
	return output
}