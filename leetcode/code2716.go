package main

import "fmt"

func main() {
	fmt.Println(minimizedStringLength("dddaaa"))
}

func minimizedStringLength(s string) int {
    characters := make(map[byte]int)
	for i :=0; i<len(s); i++{
		characters[s[i]]++
	}
	
	return len(characters)
}