package main

import "fmt"

func main() {
	fmt.Println(customSortString("cbafg", "abcd"))
}

func customSortString(order string, s string) string {
	var output string

	countMap := make(map[rune]int)
    for _, ch := range s {
        countMap[ch]++
    }

	for _, ch := range order {
        for countMap[ch] > 0 {
            output += string(ch)
            countMap[ch]--
        }
    }


	for _, ch := range s {
        if countMap[ch] > 0 {
            output += string(ch)
            countMap[ch]--
        }
    }
    
	return output
}