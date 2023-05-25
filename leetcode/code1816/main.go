package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(truncateSentence("Hello how are you Contestant", 4))
}

func truncateSentence(s string, k int) string {
	var numOfSpaces int 
	var output string
	for i:=0; i<len(s); i++{
		output += string(s[i])
		if string(s[i]) == " "{
			numOfSpaces++
			if numOfSpaces == k {
				break
			}
		}
	}
	output = strings.TrimSuffix(output, " ")
    return output
}