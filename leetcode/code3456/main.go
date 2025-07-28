package main

import "fmt"

func main() {
	fmt.Println(hasSpecialSubstring("aasdfaaa", 3))
}

func hasSpecialSubstring(s string, k int) bool {
	for startIndex := 0; startIndex < len(s); startIndex++ {
		for endIndex := startIndex + 1; endIndex <= len(s); endIndex++ {
			subarray := s[startIndex:endIndex]
			if len(subarray) == k {
				frequencyMap := make(map[byte]int)
				for i := 0; i < len(subarray); i++ {
					frequencyMap[subarray[i]]++
				}
				if len(frequencyMap) == 1 {
					return true
				}
			}
		}
	}

	return false
}
