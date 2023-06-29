package main

import (
	"strconv"
)

func main() {
	isPalindrome(121)
}

func isPalindrome(x int) bool {
	// var reversed int
	var xReversedStr string
	xStr := strconv.Itoa(x)
    for i:=len(xStr)-1;i>=0;i--{
		xReversedStr += string(xStr[i])
	}
	if xStr == xReversedStr{
		return true
	}
	return false
}