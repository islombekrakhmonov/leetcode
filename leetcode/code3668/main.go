package main

import "fmt"

func main() {
	fmt.Println(recoverOrder([]int{3, 1, 2, 5, 4}, []int{1, 3, 4}))
}

func recoverOrder(order []int, friends []int) []int {
	var output []int

	friendSet := make(map[int]struct{}, len(friends))
	for _, f := range friends {
		friendSet[f] = struct{}{}
	}

	for _, id := range order {
		if _, ok := friendSet[id]; ok {
			output = append(output, id)
		}
	}

	return output
}
