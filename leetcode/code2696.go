package main

import "fmt"

func main() {
	fmt.Println(minLength("ABFCACDB"))
}

func minLength(s string) int {
	found := true
	for found {
		found = false 
		for i:=0; i<len(s)-1; i++{
			if (s[i] == 'A' && s[i+1] == 'B') || (s[i] == 'C' && s[i+1] == 'D') {
				s = s[:i] + s[i+2:]
				found = true
				break
			}
		}
	}

    

	return len(s)
}