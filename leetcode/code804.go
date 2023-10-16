package main

import "fmt"

func main() {
	fmt.Println(uniqueMorseRepresentations([]string{"gin","zen","gig","msg"}))
}

func uniqueMorseRepresentations(words []string) int {

	uniqueStrings := make(map[string]bool)
	codes := []string{}
	for _,v := range words{
		code := ""
		for _,k := range v {
			code += morse(string(k))
		}
		codes = append(codes, code)
	}
	for _, str := range codes {
        uniqueStrings[str] = true
    }

	return len(uniqueStrings)
}

func morse(s string) (code string) {
	morseCode := []string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	morseAlphabet := make(map[string]string)

	for i,v := range alphabet{
		morseAlphabet[string(v)] = morseCode[i]
	}

	for _,v := range s {
		code += morseAlphabet[string(v)]
	}
	
	return 
}