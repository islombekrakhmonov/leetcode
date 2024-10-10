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
			if isEqual(substring) {
				outputInt++
			}
		}
	}

	return outputInt
}

func isEqual(s string) bool {
	var count0, count1 int

	for _, v := range s {
		if v == '1' {
			count1++
		} else {
			count0++
		}
	}

	return count0 == count1
}
