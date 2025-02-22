package main

import "fmt"

func main() {
	fmt.Println(numberOfLines([]int{3, 4, 10, 4, 8, 7, 3, 3, 4, 9, 8, 2, 9, 6, 2, 8, 4, 9, 9, 10, 2, 4, 9, 10, 8, 2}, "mqblbtpvicqhbrejb"))
}

func numberOfLines(widths []int, s string) []int {
	var lines, lastLineWidth int

	var alphabet = "abcdefghijklmnopqrstuvwxyz"

	mapped := make(map[byte]int)

	for i, v := range widths {
		mapped[alphabet[i]] = v
	}

	for i := 0; i < len(s); i++ {
		width := mapped[s[i]]
		if lastLineWidth+width > 100 {
			lines++
			lastLineWidth = width
		} else {
			lastLineWidth += width
		}
	}

	lines++

	return []int{lines, lastLineWidth}
}
