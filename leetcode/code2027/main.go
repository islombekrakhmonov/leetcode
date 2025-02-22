package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimumMoves("XXOX"))
}

func minimumMoves(s string) int {
	var output int

	for i := 0; i < len(s); i++ {
		if s[i] == 'X' {
			output += 1
			i += 2
		}
	}
	return output
}
