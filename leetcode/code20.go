package main

import "fmt"

func main() {
	fmt.Println(isValid("()))"))	
}

func isValid(s string) bool {

	if len(s) == 0 || len(s) % 2 != 0 {
		return false
	}

    if s[0] == ']' || s[0] == ')' || s[0] == '}' {
		return false 
	}

	var stack []byte

	for i := 0; i < len(s); i++ {
		if s[i] == '[' || s[i] == '(' || s[i] == '{' {
			stack = append(stack, s[i])
		} else if len(stack) == 0 {
			return false 
		} else {
			top := stack[len(stack)-1]
			if (top == '(' && s[i] == ')') || (top == '[' && s[i] == ']') || (top == '{' && s[i] == '}') {
				stack = stack[:len(stack)-1]
			} else {
				return false 
			}
		}
	}
	
	return len(stack) == 0
}