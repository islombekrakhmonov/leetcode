package main

import (
	"fmt"
)

func main() {
	fmt.Println(letterCombinations("23"))
}

func letterCombinations(digits string) (output []string) {
	words := make(map[string][]string)
	letters := []string{"a","b","c","d","e","f","g","h","i", "j","k","l","m", "n", "o","p","q","r","s","t","u","v","w","x","y","z"}
	counter := 2
	word := []string{}
	for i,v := range letters {
		word = append(word, v)
		
		if (i+1) % 3==0{
			key := fmt.Sprintf("%d", counter)
			if counter == 7 || counter ==9 {
				words[key] = append(word, letters[i-3]) 
			} else {
				words[key] = word
			}
			word = []string{}
			counter++
		}
	}

	

	return
}