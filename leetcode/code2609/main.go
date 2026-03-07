package main

import "fmt"

func main() {
	fmt.Println(findTheLongestBalancedSubstring("01000111"))
	fmt.Println(isBalanced("000111"))
}

func findTheLongestBalancedSubstring(s string) int {
	max := 0

	for startIndex := 0; startIndex < len(s); startIndex++ {
		for endIndex := startIndex + 1; endIndex <= len(s); endIndex++ {
			subs := endIndex - startIndex
			if subs%2 == 0 && subs > 1 {
				subarray := s[startIndex:endIndex]
				if isBalanced(subarray) {
					if subs > max {
						max = subs
					}
				}
			}
		}
	}

	return max
}

func isBalanced(s string) bool {
	if s[0] == '1' {
		return false
	}

	count0, count1 := 0, 0
	seenOne := false

	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			if seenOne {
				return false
			}
			count0++
		} else {
			count1++
			seenOne = true
		}
	}

	return count0 == count1
}
