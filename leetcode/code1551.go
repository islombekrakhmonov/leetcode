package main

import "fmt"

func main() {
	fmt.Println(minOperations(3))
}

func minOperations(n int) int {
	mid := n / 2
	operations := 0


	if n%2 == 0 {
		operations = mid * mid
	} else {
		operations = mid * (mid + 1)
	}

	return operations
}