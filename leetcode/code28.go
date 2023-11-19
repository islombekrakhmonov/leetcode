package main

import "fmt"

func main() {
	fmt.Println(strStr("leetcode","leeto"))
}

func strStr(haystack string, needle string) int {

    i, j := 0, 0

	for j < len(haystack) {
		if haystack[i] == needle[j] {
            j++
            if j == len(needle) {
                return i - j + 1
            }
        } else {
            i = i - j + 1
            j = 0
        }

        i++
	}

	return -1
}