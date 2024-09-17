package main

import "fmt"

func main() {
	fmt.Println(distributeCandies([]int{1, 1, 2, 2, 3, 3}))
}

func distributeCandies(candyType []int) int {
	max := len(candyType) / 2

	typesOfCandies := make(map[int]int)

	for _, v := range candyType {
		typesOfCandies[v]++
	}

	if len(typesOfCandies) < max {
		return len(typesOfCandies)
	}

	return max
}
