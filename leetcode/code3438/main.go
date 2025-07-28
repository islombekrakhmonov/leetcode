package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(findValidPair("2523533"))
}

func findValidPair(s string) string {

	var countMap = make(map[byte]int)
	var pairs [][]byte

	for i := 0; i < len(s)-1; i++ {
		if s[i] != s[i+1] {
			pair := []byte{s[i], s[i+1]}
			pairs = append(pairs, pair)
		}
		countMap[s[i]]++
	}

	countMap[s[len(s)-1]]++

	for _, pair := range pairs {
		num1, _ := strconv.Atoi(string(pair[0]))
		num2, _ := strconv.Atoi(string(pair[1]))
		if countMap[pair[0]] == num1 && countMap[pair[1]] == num2 {
			return string(pair)
		}
	}

	return ""
}
