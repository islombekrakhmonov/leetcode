package main

import "fmt"

func main() {
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}

func findDisappearedNumbers(nums []int) []int {
	distincts := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		distincts[nums[i]] = nums[i]
	}
	res := []int{}
	fmt.Println(distincts)
	for i := 1; i < len(distincts); i++ {
		if distincts[i] == 0 {
			res = append(res, i)
		}
	}
	return res
}
