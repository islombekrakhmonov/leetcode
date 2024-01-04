package main

import "fmt"

func main() {
	fmt.Println(groupThePeople([]int{3,3,3,3,3,1,3}))
}

func groupThePeople(groupSizes []int) [][]int {
	var output [][]int
	groups := make(map[int][]int)
	for i,size := range groupSizes {
		groups[size] = append(groups[size], i)
		if len(groups[size]) == size {
			fmt.Println(groups[size], size)
			output = append(output, groups[size])
			groups[size] = []int{}
		}
	}

	return output
}