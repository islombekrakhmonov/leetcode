package main

import (
	"fmt"
)

func main() {
	fmt.Println(canConstruct("aa", "aab"))
}

func canConstruct(ransomNote string, magazine string) bool {
	hashMap := make(map[rune]int)
	for _, v := range magazine {
		hashMap[v]++
	}

	for _, v := range ransomNote {
		hashMap[v]--
		if hashMap[v] < 0 {
			return false
		}
	}

	return true
}
