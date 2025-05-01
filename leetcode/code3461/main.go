package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(hasSameDigits("3902"))
}

func hasSameDigits(s string) bool {
	for len(s) != 2 {
		newS := ""
		for i := 0; i < len(s)-1; i++ {
			a, _ := strconv.Atoi(string(s[i]))
			b, _ := strconv.Atoi(string(s[i+1]))
			sum := strconv.Itoa((a + b) % 10)
			newS += sum
		}
		s = newS
	}

	return s[1] == s[0]
}

//(s[0] + s[1]) % 10 = (3 + 9) % 10 = 2
