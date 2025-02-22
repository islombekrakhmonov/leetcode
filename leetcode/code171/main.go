package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(titleToNumber("YZ")) // 1
}

// YZ = 25*(26^1) + 26*(26^0)
// XYZ = 24*(26^2) + 25*(26^1) + 26*(26^0)

func titleToNumber(columnTitle string) int {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	alphabetMap := make(map[rune]int)

	for i, v := range alphabet {
		alphabetMap[v] = i + 1
	}

	titleNum := 0

	l := len(columnTitle) - 1

	for _, column := range columnTitle {
		multiplayer := math.Pow(26, float64(l))
		columnNum := alphabetMap[column] * int(multiplayer)
		titleNum += columnNum
		l--
	}

	return titleNum
}
