package main

import "fmt"

func main() {
	fmt.Println(divideArray([]int{3, 2, 3, 2, 2, 2}))
}

func divideArray(nums []int) bool {

	hashMap := make(map[int]int)

	for _, num := range nums {
		hashMap[num]++
	}

	for _, count := range hashMap {
		if count%2 != 0 {
			return false
		}
	}

	return true
}
