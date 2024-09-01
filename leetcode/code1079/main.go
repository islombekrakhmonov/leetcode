package main

import "fmt"

func main() {

	fmt.Println(numTilePossibilities("AAC"))
}

func numTilePossibilities(tiles string) int {
	letterCount := make(map[rune]int)
	for _, ch := range tiles {
		letterCount[ch]++
	}

	return backtrack(letterCount)
}

func backtrack(letterCount map[rune]int) int {
	count := 0

	for ch, cnt := range letterCount {
		if cnt > 0 {
			// Use this letter
			letterCount[ch]--
			count += 1 + backtrack(letterCount) // count this sequence and all following sequences

			// Backtrack
			letterCount[ch]++
		}
	}

	return count
}
