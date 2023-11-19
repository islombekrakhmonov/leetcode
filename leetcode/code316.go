package main

import (
	"fmt"
)

func main() {
	fmt.Println(removeDuplicateLetters("cbacdcbc"))

}

func removeDuplicateLetters(s string) string {
	letters := make(map[string]int)
	var singled []string
	var output string

	for _, v := range s {
		letters[string(v)]++
		fmt.Println(byte(v))
	}

	for key, _ := range letters {
		singled = append(singled, key)

	}

	for i:=0; i<len(singled)-1; i++{
		for j:=0; j<len(singled)-i-1; j++{
			if singled[j] > singled[j+1] {
				singled[j], singled[j+1] = singled[j+1], singled[j]
			}
		}
	}

	for _,v := range singled {
		output += v
	}


	return output
}