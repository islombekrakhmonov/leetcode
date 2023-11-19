package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverseOnlyLetters("a-bC-dEf-ghIj"))
}

func reverseOnlyLetters(s string) string {
	output := []byte(s)
	// pos := 0


	j := len(s)-1
	i:=0

	for i<j {
		for i<j && !isEnglishLetter(output[i]){
			i++
		}
		for i<j && !isEnglishLetter(output[j]) {
			j--
		}
			output[i], output[j] = output[j], output[i] 
			i++
			j--
		}
	return string(output)

	// for i, c := range s {
	// 	if !isEnglishLetter(c) {
	// 		output[i] = c
	// 	} else {
	// 		for !isEnglishLetter(rune(s[len(s)-1-pos])) {
	// 			pos++
	// 		}
	// 		output[i] = rune(s[len(s)-1-pos])
	// 		pos++
	// 	}
	// }

}

func isEnglishLetter(char byte) bool {
	return (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z')
}
