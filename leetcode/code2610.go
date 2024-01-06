package main

import "fmt"

func main() {
	fmt.Println(findMatrix([]int{1,3,4,1,2,3,1}))
}

func findMatrix(nums []int) [][]int {
    var output [][]int
	usedIndices := make(map[int]bool)

	for _, v := range nums {
		if !usedIndices[v] {
			usedIndices[v] = true
			group := []int{v}
			
			// Include all subsequent non-usedIndices values in the current group
			for _, v2 := range nums {
				if !usedIndices[v2] {
					group = append(group, v2)
					usedIndices[v2] = true
				}
			}
			
			output = append(output, group)
		}
	}


	return output
}