package main

import "fmt"

func main() {
	fmt.Println(reverseByType(")ebc#da@f("))
}

func reverseByType(s string) string {
	letters, spec := "", ""

	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			letters += string(char)
		} else {
			spec += string(char)
		}
	}

	fmt.Println(letters)
	letters = string(reverseSlice([]rune(letters)))
	fmt.Println(letters)
	spec = string(reverseSlice([]rune(spec)))

	var output string
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			output += string(letters[0])
			if len(letters) == 1 {
				letters = ""
			} else {
				letters = letters[1:]
			}
		} else {
			output += string(spec[0])
			if len(spec) == 1 {
				spec = ""
			} else {
				spec = spec[1:]
			}
		}
	}

	return output
}

func reverseSlice(s []rune) []rune {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
	return s
}
