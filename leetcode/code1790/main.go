package main

import "fmt"

func main() {
	fmt.Println(areAlmostEqual("bank", "kanb"))
}

func areAlmostEqual(s1 string, s2 string) bool {
	var notEqualChars int
	var indexes1 []int
	var indexes2 []int

	if s1 == s2 {
		return true
	}

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			indexes1 = append(indexes1, i)
			indexes2 = append(indexes2, i)
			notEqualChars++
		}
	}

	if notEqualChars != 2 && notEqualChars != 0 {
		return false
	}

	equal := false

	originalS1 := s1
	s1Runes := []rune(s1)
	s1Runes[indexes1[0]], s1Runes[indexes1[1]] = s1Runes[indexes1[1]], s1Runes[indexes1[0]]
	s1 = string(s1Runes)
	if s1 == s2 {
		equal = true
	}

	s1 = originalS1
	s2Runes := []rune(s2)
	s2Runes[indexes2[0]], s2Runes[indexes2[1]] = s2Runes[indexes2[1]], s2Runes[indexes2[0]]
	s2 = string(s2Runes)
	if s1 == s2 {
		equal = true
	}

	return equal
}
