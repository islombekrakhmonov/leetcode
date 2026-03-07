package main

import "fmt"

func main() {
	fmt.Println(totalNumbers([]int{1, 2, 3, 4}))
}

func totalNumbers(digits []int) int {
	l := len(digits)

	return factorial(l) / factorial(l-3)
}

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}
