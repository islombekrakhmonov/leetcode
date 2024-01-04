package main

import "fmt"

func main() {
	fmt.Println(numEquivDominoPairs([][]int{{1,2},{1,2},{1,1},{1,2}, {2,2}}))
}

func numEquivDominoPairs(dominoes [][]int) int {
	output := 0
	countMap := make(map[[2]int]int)

	for _, domino := range dominoes {
        if domino[0] > domino[1] {
            domino[0], domino[1] = domino[1], domino[0]
        }

        dominoKey := [2]int{domino[0], domino[1]}

        countMap[dominoKey]++

        output += countMap[dominoKey] - 1
    }

    return output
}