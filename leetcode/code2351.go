package main

import "fmt"

func main() {
	fmt.Println(repeatedCharacter("abacddbec"))
}

func repeatedCharacter(s string) byte {
	characters := make(map[byte]int)
	for i := 0; i<len(s);i++{
		if _, exists := characters[s[i]]; exists{ 
			return s[i]
		} else {
			characters[s[i]]++
		}
	}
	
    
	return ('l')
}