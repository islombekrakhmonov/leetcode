package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(numDifferentIntegers("4"))
}

func numDifferentIntegers(word string) int {
    wordRunes := []rune(word)
    for i := 0; i < len(wordRunes); i++ {
        if (wordRunes[i] >= 'a' && wordRunes[i] <= 'z') || (i > 0 && wordRunes[i-1] == ' ' && wordRunes[i] == '0' && i+1 < len(wordRunes) && unicode.IsDigit(wordRunes[i+1])) || i == 0 && i+1 < len(wordRunes) && unicode.IsDigit(wordRunes[i+1]){
            wordRunes[i] = ' '
        }
    }
    count := make(map[string]int)

    splitted := strings.Split(string(wordRunes), " ")
    for _, v := range splitted {
        if v != "" {
            count[v]++
        }
    }

    return len(count)
}
