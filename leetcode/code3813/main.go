package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(vowelConsonantScore("i3"))
}

func vowelConsonantScore(s string) int {
	v, c := 0, 0

	for _, char := range s {
		if isVowel(char) {
			v++
		} else if char >= 'a' && char <= 'z' {
			c++
		}
	}

	if v == 0 || c == 0 {
		return 0
	}

	return int(math.Floor(float64(v / c)))
}

func isVowel(char rune) bool {
	vowelChar := []rune{'a', 'e', 'i', 'o', 'u'}
	for _, v := range vowelChar {
		if char == v {
			return true
		}
	}
	return false
}
