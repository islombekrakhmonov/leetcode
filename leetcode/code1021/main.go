package main

import "fmt"

func main() {

	fmt.Println(removeOuterParentheses("(()())(())")) // Output: ()()()
}

func removeOuterParentheses(s string) string {
	var result []rune
	depth := 0

	for _, char := range s {
		if char == '(' {
			if depth > 0 {
				result = append(result, char) // Add '(' if it's not the outermost
			}
			depth++
		} else if char == ')' {
			depth--
			if depth > 0 {
				result = append(result, char) // Add ')' if it's not the outermost
			}
		}
	}

	return string(result)
}
