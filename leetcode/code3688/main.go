package main

import "fmt"

func main() {
	fmt.Println(evenNumberBitwiseORs([]int{1, 2, 3}))
}

func evenNumberBitwiseORs(nums []int) int {
	var output int

	for _, num := range nums {
		if num%2 == 0 {
			output |= num
		}
	}

	return output
}
