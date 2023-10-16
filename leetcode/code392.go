package main

import "fmt"

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc"))
}

func isSubsequence(s string, t string) bool {
    var sIndex, tIndex int

	for sIndex < len(s) && tIndex < len(t) {
        if s[sIndex] == t[tIndex] {
            sIndex++
        }
        tIndex++
    }

	return sIndex == len(s)
}