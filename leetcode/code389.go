package main

import "fmt"

func main() {
	fmt.Println(findTheDifference("abcd", "abcdy"))
}

func findTheDifference(s string, t string) byte {
    word1Mapped := make(map[byte]int)
	word2Mapped := make(map[byte]int)

	for i := 0; i<len(s); i++{
		word1Mapped[(s[i])]++
	}

	for i := 0; i<len(t); i++{
		word2Mapped[(t[i])]++
	} 


	for i, v := range word1Mapped{
		if v - word2Mapped[i] > 0 {
			return i
		}
	}

	for i, v := range word2Mapped{
		if v - word1Mapped[i] > 0 {
			return i
		}
	}

	return ' '
}