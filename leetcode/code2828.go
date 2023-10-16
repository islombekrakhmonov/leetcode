package main

import "fmt"

func main() {
	fmt.Println(isAcronym([]string{"alice","bob","charlie"},"abc"))
}

func isAcronym(words []string, s string) bool {
	var letters string 
	for _,v := range words{
		letters += string(v[0])
	}

	return letters == s 
}