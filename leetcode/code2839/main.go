package main

import "fmt"

func main() {
	fmt.Println(canBeEqual("bnxw", "bwxn"))
}

func canBeEqual(s1 string, s2 string) bool {

	if s1 == s2 {
		return true
	}

	var indexes1 []int
	var indexes2 []int

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			indexes1 = append(indexes1, i)
			indexes2 = append(indexes2, i)
		}
	}

	equal := false

	originalS1 := s1
	s1Runes := []rune(s1)

	for i := 0; i < len(indexes1)-1; i++ {
		for j := i + 1; j < len(indexes1); j++ {
			if indexes1[j]-indexes1[i] == 2 {
				s1Runes[indexes1[i]], s1Runes[indexes1[j]] = s1Runes[indexes1[j]], s1Runes[indexes1[i]]
				s1 = string(s1Runes)
				if s1 == s2 {
					equal = true
				}
			}
		}
	}

	s1 = originalS1
	s2Runes := []rune(s2)

	for i := 0; i < len(indexes2)-1; i++ {
		for j := i + 1; j < len(indexes2); j++ {
			if j-i == 2 {
				s2Runes[indexes2[i]], s2Runes[indexes2[j]] = s2Runes[indexes2[j]], s2Runes[indexes2[i]]
				s2 = string(s2Runes)
				if s1 == s2 {
					equal = true
				}
			}
		}
	}

	return equal
}
