package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isCircularSentence("leetcode exercises sound delightful"))
}

func isCircularSentence(sentence string) bool {

	splitted := strings.Split(sentence, " ")

	for i := 0; i < len(splitted)-1; i++ {
		if splitted[i][len(splitted[i])-1] != splitted[i+1][0] {
			return false
		}
	}

	if sentence[0] != sentence[len(sentence)-1] {
		return false
	}

	return true
}

// if len(splitted) == 1 {
//
// }
