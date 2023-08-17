package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isStrictlyPalindromic(9))
}

func isStrictlyPalindromic(n int) bool {
	remainder := 0
	var output string
	for i:=2; i<=n-2; i++{
		number := n
		output = ""
		remainder = 0
		for number != 0{
			remainder = number % i
			number /= i
			output += strconv.Itoa(remainder)
			isPalindrome := isPalindromic(output)
			if !isPalindrome {
				return false
			}
		}
	}
    
	return true
}

func isPalindromic(number string) bool {
	var numberReversed string
	for i:=len(number)-1;i>=0;i--{
		numberReversed += string(number[i])
	}
	return number == numberReversed
}
