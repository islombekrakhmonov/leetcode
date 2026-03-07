package main

import "fmt"

func main() {
	fmt.Println(validStrings(1))
}

func validStrings(n int) []string {
	var output []string

	for i := 0; i < (1 << n); i++ {
		binary := fmt.Sprintf("%0*b", n, i)
		if isAdjacent(binary) {
			continue
		}

		output = append(output, binary)
	}

	return output
}

func isAdjacent(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == '0' && s[i+1] == '0' {
			return true
		}
	}

	return false
}
