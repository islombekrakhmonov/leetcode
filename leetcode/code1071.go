package main

import (
	"fmt"
)

func main() {
	fmt.Println(gcdOfStrings("ABCABC", "ABC"))
}

func gcdOfStrings(str1 string, str2 string) string {
	commonDivisor := make(map[string]int)
	for i := 0; i < len(str1); i++ {
		for j := 0; j < len(str2); j++ {
			if str1[i] == str2[j] {
				commonDivisor[string(str1[i])]++
			}
		}
	}

	var output string

	for key := range commonDivisor {
		output += key
	}

	return output
}
