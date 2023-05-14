package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "abcdefd"
	var ch byte = 'd'

	fmt.Println(reversePrefix(input,ch))
	
}

func reversePrefix(word string, ch byte)string {
	var result string
	var gross string
	var upto string
	count := strings.Count(word, string(ch))
	for i,v := range word{
		if byte(v) == ch{
			upto = word[:i+1]
			gross = word[i+1:]
			break
		} else if count == 0{
			return word
		}
	}
    for _,v := range upto {
	    result = string(v) + result
	}
	result +=gross
   return result
}