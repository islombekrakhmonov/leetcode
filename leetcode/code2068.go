package main

import "fmt"

func main() {
	fmt.Println(checkAlmostEquivalent("aaaa", "bccb"))
}

func checkAlmostEquivalent(word1 string, word2 string) bool {
	word1Mapped := make(map[byte]int)
	word2Mapped := make(map[byte]int)

	for i := 0; i<len(word1); i++{
		word1Mapped[(word1[i])]++
	}

	for i := 0; i<len(word2); i++{
		word2Mapped[(word2[i])]++
	} 


	for i, v := range word1Mapped{
		fmt.Println(string(i), v)
		if v - word2Mapped[i] > 3 {
			return false 
		}
	}

	for i, v := range word2Mapped{
		fmt.Println(string(i), v)
		if v - word1Mapped[i] > 3 {
			return false 
		}
	}

	return true
}