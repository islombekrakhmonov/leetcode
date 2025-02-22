package main

import "fmt"

func main() {
	fmt.Println(findChampion(3, [][]int{{0, 1}, {1, 2}}))
}

func findChampion(n int, edges [][]int) int {
	sumScore := make(map[int]int)

	for n > 0 {
		sumScore[n-1] = 0
		n--
	}

	for i := 0; i < len(edges); i++ {
		sumScore[edges[i][1]]++
	}

	var chempions []int

	for i, v := range sumScore {
		if v == 0 {
			chempions = append(chempions, i)
		}
	}

	if len(chempions) > 1 {
		return -1
	}

	return chempions[0]
}
