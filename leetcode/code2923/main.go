package main

import "fmt"

func main() {
	fmt.Println(findChampion([][]int{{0, 0, 1}, {1, 0, 1}, {0, 0, 0}}))
}

func findChampion(grid [][]int) int {

	sumScore := make(map[int]int)

	for i := 0; i < len(grid); i++ {
		var sum int
		for j := 0; j < len(grid[i]); j++ {
			if i != j {
				sum += grid[i][j]
			}
		}
		sumScore[i] = sum
	}

	var max, maxI int

	for i, v := range sumScore {
		if v > max {
			max = v
			maxI = i
		}
	}

	return maxI
}
