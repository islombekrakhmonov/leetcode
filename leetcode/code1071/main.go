package main

import "fmt"

func main() {

	fmt.Println(gcdOfStrings("ABABAB", "ABAB")) // ABC
}

func gcdOfStrings(str1 string, str2 string) string {

	if str1+str2 != str2+str1 {
		return ""
	}

	output := gcd(len(str1), len(str2))

	return str1[:output]
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
