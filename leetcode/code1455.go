package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isPrefixOfWord("i love eating burger", "burg"))
}

func isPrefixOfWord(sentence string, searchWord string) int {
	splitted := strings.Split(sentence, " ")
    for i:=0; i<len(splitted); i++{
		if len(splitted[i]) >= len(searchWord) {
			if searchWord == splitted[i][:len(searchWord)] {
				return i+1
			}
		}
	}

	return -1
}