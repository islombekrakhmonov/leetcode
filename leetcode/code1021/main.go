package main

import "fmt"

func main() {

	fmt.Println(removeOuterParentheses("(()())(())"))
}

func removeOuterParentheses(s string) string {
	var result []rune
	charCount := make(map[rune]int)

	first := true
	for _, char := range s {
		if first {
			first = false
			charCount['(']++
			continue
		}

		if char == '(' {
			charCount['(']++
		} else {
			charCount[')']++
		}

		if charCount['('] == charCount[')'] {
			first = true
			charCount = make(map[rune]int)
			continue
		}

		result = append(result, char)
	}

	return string(result)
}
