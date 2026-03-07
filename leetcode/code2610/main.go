package main

import "fmt"

func main() {
	fmt.Println(findMatrix([]int{1, 3, 4, 1, 2, 3, 1}))
}

func findMatrix(nums []int) [][]int {
	var output [][]int
	frequencyMap := make(map[int]int)

	for _, num := range nums {
		frequencyMap[num]++
	}

	for len(frequencyMap) > 0 {
		var group []int
		for num := range frequencyMap {
			group = append(group, num)
			frequencyMap[num]--
			if frequencyMap[num] == 0 {
				delete(frequencyMap, num)
			}
		}
		output = append(output, group)
		if len(frequencyMap) == 0 {
			break
		}
	}

	return output
}
