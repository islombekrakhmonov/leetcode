package main

import "fmt"

func main() {
	fmt.Println(maxScore("011101"))
}

func maxScore(s string) int {
	var output int
	for i := 1; i < len(s); i++ {
		var sum int
		left := s[:i]
		for _, value := range left {
			if value == '0' {
				sum++
			}
		}
		right := s[i:]
		for _, value := range right {
			if value == '1' {
				sum++
			}
		}
		if sum > output {
			output = sum
		}
	}
	return output
}
