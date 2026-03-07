package main

import "fmt"

func main() {
	fmt.Println(largestEven("1"))
}

func largestEven(s string) string {

	if s[len(s)-1]%2 == 0 {
		return s
	}

	for len(s) > 0 {
		if s[len(s)-1]%2 != 0 {
			s = s[:len(s)-1]
		} else {
			break
		}
	}

	return s
}
