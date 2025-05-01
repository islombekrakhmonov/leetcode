package main

import "fmt"

func main() {
	fmt.Println(rotateString("abcde", "cdeab"))
}

func rotateString(s string, goal string) bool {
	i := len(s)

	for i > 0 {
		s = rotateLeft(s, 1)
		if s == goal {
			return true
		}
		i--
	}

	return false
}

func rotateLeft(s string, x int) string {
	n := len(s)
	x = x % n

	return s[x:] + s[:x]
}
