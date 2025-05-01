package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumUnits([][]int{{1, 3}, {5, 5}, {2, 5}, {4, 2}, {4, 1}, {3, 1}, {2, 2}, {1, 3}, {2, 5}, {3, 2}}, 35))
}

func maximumUnits(boxTypes [][]int, truckSize int) int {
	var output int
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})

	for len(boxTypes) > 0 && truckSize > 0 {
		for i := 0; i < len(boxTypes); i++ {
			if boxTypes[i][0] > 0 {
				output += boxTypes[i][1]
				truckSize--
				boxTypes[i][0]--
				if boxTypes[i][0] == 0 {
					boxTypes = boxTypes[i:]
				}
				break
			} else {
				if len(boxTypes) == 1 && boxTypes[i][0] == 0 {
					boxTypes = nil
				}
			}
		}
	}

	return output
}
