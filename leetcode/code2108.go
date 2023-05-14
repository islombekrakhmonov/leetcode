package main

import "fmt"

func main() {
	words := []string{"abc","car","ada","racecar","cool"}
	fmt.Println(firstPalindrome(words))
	
}

func firstPalindrome(words []string) string {
	for i:= range words{
		x:=palin(words[i])
		if x == words[i]{
			return words[i]
		}
	}

	return ""
}

func palin(t string)string{
	var text string
	for x :=len(t)-1; x>=0;x--{
		text +=string(t[x])
	}

	return text

}