package main

import "fmt"

func main() {
	fmt.Println(halvesAreAlike("book"))
}

func halvesAreAlike(s string) bool {

	vowels := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'A': true, 'E': true, 'I':true, 'O': true, 'U': true}

	l := len(s) / 2
    a, b := s[:l], s[l:]

	countA, countB := 0,0

	for i:=0; i<len(a); i++ {
		if vowels[a[i]] {
			countA++
		}
	}

	for i:=0; i<len(b); i++ {
		if vowels[b[i]] {
			countB++
		}
	}

	return countA == countB 
}
