package main

import "fmt"

func main() {
	fmt.Println(scoreOfString("hello"))
}

func scoreOfString(s string) int {
	var sum int

	for i := 0; i < len(s)-1; i++ {
		diff := substraction(s[i], s[i+1])
		sum += int(diff)
	}

	return sum
}

func substraction(a byte, b byte) byte {
	if a >= b {
		return a - b
	} else {
		return b - a
	}
}
