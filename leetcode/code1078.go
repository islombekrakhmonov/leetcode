package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(findOcurrences("alice is a good girl she is a good student", "a", "good"))
}

func findOcurrences(text string, first string, second string) []string {
    splitted := strings.Split(text, " ")
	var output []string

	for i :=0; i<len(splitted)-2; i++{
		if splitted[i] == first && splitted[i+1] == second {
			output = append(output, splitted[i+2])
		}
	}
	return output
}