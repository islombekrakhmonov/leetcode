package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(divideString("bgycymgbblobvpdf", 67, 'u'))
}

func divideString(s string, k int, fill byte) []string {

	result := make([]string, 0)

	for i := 0; i < len(s); i += k {
		var subString string

		if i+k < len(s) {
			subString = s[i : i+k]
			fmt.Println(subString)
		} else {
			subString = s[i:] + strings.Repeat(string(fill), k-len(s[i:]))
			fmt.Println(subString)
		}

		result = append(result, subString)
	}

	return result
}
