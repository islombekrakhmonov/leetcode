package main

import "fmt"

func main() {
	fmt.Println(finalValueAfterOperations([]string{"--X", "X++", "X++"}))
}

func finalValueAfterOperations(operations []string) int {
	var output int

	for _, v := range operations {
		if contains(v, '-') {
			output--
		} else {
			output++
		}
	}

	return output
}

func contains(s string, b rune) bool {
	for _, value := range s {
		if value == b {
			return true
		}
	}

	return false
}
