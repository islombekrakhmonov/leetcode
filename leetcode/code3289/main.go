package main

import "fmt"

func main() {
	fmt.Println(getSneakyNumbers([]int{0, 3, 2, 1, 3, 2}))
}

func getSneakyNumbers(nums []int) []int {

	var output []int

	numCount := make(map[int]int)

	for _, v := range nums {
		numCount[v]++
	}

	for i, v := range numCount {
		if v == 2 {
			output = append(output, i)
		}
	}

	return output
}
