package main

import "fmt"

func main() {
	fmt.Println(numOfUnplacedFruits([]int{4, 2, 5}, []int{3, 5, 4}))
}

func numOfUnplacedFruits(fruits []int, baskets []int) int {

	for i := 0; i < len(fruits); i++ {
		for j := 0; j < len(baskets); j++ {
			if fruits[i] <= baskets[j] {
				baskets = append(baskets[:j], baskets[j+1:]...)
				break
			}
		}
	}

	return len(baskets)
}
