package main

import "fmt"

func main() {
	fmt.Println(stableMountains([]int{1, 2, 3, 4, 5}, 2))
}

func stableMountains(height []int, threshold int) []int {
	var returnSlice []int

	for i := 1; i < len(height); i++ {
		if height[i-1] > threshold {
			returnSlice = append(returnSlice, i)
		}
	}

	return returnSlice
}
