package main

import "fmt"

func main() {
	fmt.Println(maxContainers(2, 3, 15))
}

func maxContainers(n int, w int, maxWeight int) int {
	return min(maxWeight/w, n*n)
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}
