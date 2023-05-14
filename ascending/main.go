package main

import (
	"fmt"
	"strings"
)

func main() {
	// input := []int{4,53,14}
	sentence := "thequickbrownfoasdfghjhgfxjumpsoverthelazydog"
	// fmt.Println(ascending(input))
	fmt.Println(checkIfPangram(sentence))
}

func ascending(input []int)[]int{
	// var smth []int
	var result []int
	var output []int
	for i,v := range input{
		output = append(output, v+i+1)
	}
	for _,v := range output{
		if v < 9{
			result = append(result, v)
		} else if v > 9{
			result = append(result, v%10)
		} 
	}
	return result
}

func checkIfPangram(sentence string) bool {
	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	var pangram = true
	var count int
	var countSlice []int
	for _,t := range alphabet{
		count = strings.Count(sentence, t)
		countSlice = append(countSlice, count)
	}
	for _,v := range countSlice{
		if v == 0 {
			pangram = false 
		} 
	}
	return pangram
}
