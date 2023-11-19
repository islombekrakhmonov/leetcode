package main

import "fmt"

func main() {
	fmt.Println(firstUniqChar("aabb"))
}

func firstUniqChar(s string) int {
	charCount := make(map[rune]int)

	for _, v := range s {
		charCount[v]++
	}

	for i, ch := range s {
        if charCount[ch] == 1 {
            return i
        }
    } 
	
	return -1
}