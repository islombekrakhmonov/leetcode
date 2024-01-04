package main

import "fmt"

func main() {
	fmt.Println(checkDistances("abaccb", []int{1,3,0,5,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}))
}

func checkDistances(s string, distance []int) bool {
	for i := 0; i < len(s); i++ {
		charIndex := int(s[i] - 'a')
		expectedDistance := distance[charIndex]

		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				actualDistance := j - i - 1 
				if actualDistance != expectedDistance {
					return false
				}
				break 
			}
		}
	}

	return true
}