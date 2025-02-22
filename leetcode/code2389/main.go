package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(answerQueries([]int{736411, 184882, 914641, 37925, 214915}, []int{331244, 273144, 118983, 118252, 305688, 718089, 665450}))
}

func answerQueries(nums []int, queries []int) []int {

	sort.Ints(nums)

	answer := make([]int, len(queries))

	for i, query := range queries {
		var sum int
		var count int

		for _, num := range nums {
			if sum+num <= query {
				sum += num
				count++
			} else {
				break
			}
		}

		answer[i] = count
	}

	return answer
}
