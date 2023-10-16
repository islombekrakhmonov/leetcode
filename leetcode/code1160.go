package main

import "fmt"

func main() {
	fmt.Println(countCharacters([]string{"cat","bt","hat","tree"}, "atach"))
}

func countCharacters(words []string, chars string) int {

	var output int
	for _,v := range words {
		count := 0
		charCopy := chars
		for _, letter := range v {
			for j:=0; j<len(charCopy); j++{
				if charCopy[j] == byte(letter) {
					count++ 
					charCopy = charCopy[:j]+charCopy[j+1:]
					fmt.Println(charCopy)
					break
				}			}
			if count == len(v) {
				output += len(v)
			}
		}
	}

	return output
}