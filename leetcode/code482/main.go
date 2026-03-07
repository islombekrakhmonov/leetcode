package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(licenseKeyFormatting("5F3Z-2e-9-w", 4))
}

func licenseKeyFormatting(s string, k int) string {
	var output string

	reversed := string(reverseString([]byte(s)))

	for i := 0; i < len(reversed); i++ {
		if s[i] == '-' {
			continue
		}
		fmt.Println(string(s[i]))

		if i%k == 0 {
			output = "-" + output
		}
		fmt.Println(output)
		output = strings.ToUpper(string(s[i])) + output
	}

	reversed = string(reverseString([]byte(output)))

	if reversed[0] == '-' {
		return output[1:]
	}

	return reversed
}

func reverseString(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

/*
remove all the '-';
reverse the string
3.if some char is lower convert it to upper
4.when i%k==0 add '-' else keep on adding normally to the string ;
5.at the end reverse the string you built and return.
*/
