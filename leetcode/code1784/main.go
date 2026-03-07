package main

import "fmt"

func main() {
	fmt.Println(checkOnesSegment("011"))
}

func checkOnesSegment(s string) bool {
	if len(s) == 1 {
		if s[0] == '1' {
			return true
		} else {
			return false
		}
	}

	firstChar := s[0]
	seenZero, seenOne := false, false

	if firstChar == '0' {
		for i := 1; i < len(s); i++ {
			if s[i] == '1' {
				seenOne = true
			} else if seenOne {
				return false
			}
		}
	} else {
		for i := 1; i < len(s); i++ {
			if s[i] == '0' {
				seenZero = true
			} else if seenZero && s[i] == '1' {
				return false
			}
		}
	}

	return true
}
