package main

import "fmt"

func main() {

	fmt.Println(minimumOperations([]int{1, 2, 3, 4, 5, 6}))
}

func minimumOperations(nums []int) int {
	var output int

	for _, v := range nums {
		if (v+1)%3 == 0 || (v-1)%3 == 0 {
			output++
		}
	}

	return output
}
