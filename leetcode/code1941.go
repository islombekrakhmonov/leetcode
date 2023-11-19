package main

import "fmt"

func main() {
	fmt.Println(areOccurrencesEqual("abacbc"))
}

func areOccurrencesEqual(s string) bool {
	mapped := make(map[rune]int)
	
	for _, v := range s {
		mapped[v]++
	}

	for _,v := range mapped {
		for _, k := range mapped {
			if v != k {
				return false 
			}
		}
	}

	return true
}