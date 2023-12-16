package main

import "fmt"

func main() {
	fmt.Println(buddyStrings("ab", "ba"))
	fmt.Println(swapAdjacentLetters("golang"))
}

func buddyStrings(s string, goal string) bool {

	if len(s) != len(goal) {
        return false
    }


	var index int 
	for i :=1; i<len(s); i++ {
		swappedString := s[:index] + string(s[i]) + s[:i] + string(s[index]) + s[i:]
		if swappedString == goal {
			return true
		}
		index++
	}
    
	return false 
}

func swapAdjacentLetters(s string) []string {
	var output []string

	for i := 0; i < len(s)-1; i++ {
		swappedString := s[:i] + string(s[i+1]) + s[i:i+1] + string(s[i]) + s[i+2:]
		output = append(output, swappedString)
	}

	return output
}