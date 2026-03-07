package main

import "fmt"

func main() {
	// fmt.Println(isPrefixString("iloveleetcode", []string{"i", "love", "leetcode", "apples"}))
	// fmt.Println(isPrefixString("ileetcode", []string{"i", "love", "leetcode", "apples"}))
	fmt.Println(isPrefixString("aaa", []string{"aa", "aaa", "fjaklfj"}))
}

func isPrefixString(s string, words []string) bool {

	if s[0] != words[0][0] {
		return false
	}

	if s == words[0] {
		return true
	}

	for i := 0; i < 1; i++ {
		if len(s) < len(words[i]) {
			return false
		}
		concat := words[i]
		if s == concat {
			return true
		}

		for j := 1; j < len(words); j++ {
			concat += words[j]
			if s == concat {
				return true
			}
		}
	}

	return false
}
