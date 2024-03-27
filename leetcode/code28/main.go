package main

import "fmt"

func main() {
	fmt.Println(strStr("hello", "ll"))
}

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] {
			if len(haystack)-i < len(needle) {
				return -1
			}
			if haystack[i:i+len(needle)] == needle {
				return i
			}
		}
	}

	return -1
}
