package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(freqAlphabets("10#11#12"))
}

func freqAlphabets(s string) string {

	alphabet := "abcdefghijklmnopqrstuvwxyz"
	var output string


	for i:=0; i<len(s); i++{
		if i+2 < len(s) && (s[i] == '1' || s[i] == '2') && s[i+2] == '#' {
			index := string(s[i])+string(s[i+1])
			indexInt, _ := strconv.Atoi(index)
			output += string(alphabet[indexInt-1])
			i+=2
		} else {
			indexInt, _ := strconv.Atoi(string(s[i]))
			output += string(alphabet[indexInt-1])
		}
	}
	return output
}