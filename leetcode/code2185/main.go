package main

import "fmt"

func main() {
	fmt.Println(prefixCount([]string{"pay","attention","practice","attend"}, "at"))
}

func prefixCount(words []string, pref string) int {
	var output int 
    for _, word := range words {
		if len(word) >= len(pref) && word[:len(pref)] == pref {
			output++
		}
	}
	return output 
}