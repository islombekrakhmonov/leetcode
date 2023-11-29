package main

import "fmt"

func main() {
	fmt.Println(distinctIntegers(5))
}

func distinctIntegers(n int) int {
	return max(1, n-1)
}

func max(i, n int) int {
	if i < n {
		return n
	} else {
		return i
	}
}