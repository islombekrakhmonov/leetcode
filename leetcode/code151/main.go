package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseWords("a good   example"))
}

func reverseWords(s string) string {

	splitted := strings.Split(s, " ")

	s = ""
	for i := len(splitted) - 1; i >= 0; i-- {
		if splitted[i] == "" {
			continue
		}
		s += splitted[i] + " "
	}

	s = strings.TrimSuffix(s, " ")

	return s
}
