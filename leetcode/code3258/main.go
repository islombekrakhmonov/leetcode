package main

import "fmt"

func main() {
	fmt.Println(countKConstraintSubstrings("10101", 1))
}

func countKConstraintSubstrings(s string, k int) int {

	var output int

	for startIndex := 0; startIndex < len(s); startIndex++ {
		for endIndex := startIndex + 1; endIndex <= len(s); endIndex++ {
			substring := s[startIndex:endIndex]
			zeros, ones := num1and0s(substring)
			if zeros <= k || ones <= k {
				output++
			}
		}
	}

	return output
}

func num1and0s(s string) (zeros, ones int) {

	for _, num := range s {
		if num == '1' {
			ones++
		} else {
			zeros++
		}
	}

	return
}
