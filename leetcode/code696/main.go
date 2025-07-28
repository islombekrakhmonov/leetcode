package main

import "fmt"

func main() {
	fmt.Println(countBinarySubstrings("00110011"))
}

func countBinarySubstrings(s string) int {

	var outputInt int

	for startIndex := 0; startIndex < len(s); startIndex++ {
		for endIndex := startIndex + 1; endIndex <= len(s); endIndex++ {
			substring := s[startIndex:endIndex]
			if isEqualAndConsecutive(substring) {
				outputInt++
			}
		}
	}

	return outputInt
}

func isEqualAndConsecutive(s string) bool {
	if len(s) == 0 {
		return false
	}

	var count0, count1 int
	var transitions int
	prev := s[0]

	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			count0++
		} else if s[i] == '1' {
			count1++
		}

		if i > 0 && s[i] != prev {
			transitions++
			prev = s[i]
		}
	}

	return transitions == 1 && count0 == count1
}
