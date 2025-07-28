package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maximumPopulation([][]int{{1900, 1950}, {1910, 1940}, {1920, 1935}}))
}

func maximumPopulation(logs [][]int) int {
	min := math.MaxInt
	max := math.MinInt
	maxCount := 0
	minYear := 0

	for _, log := range logs {
		if log[0] < min {
			min = log[0]
		}
		if log[1] > max {
			max = log[1]
		}
	}

	yearsMap := make(map[int]int)

	for i := min; i <= max; i++ {
		yearsMap[i] = 0
	}

	for year, num := range yearsMap {
		for _, log := range logs {
			if year >= log[0] && year < log[1] {
				num++
			}
		}
		yearsMap[year] = num
	}

	for year, num := range yearsMap {
		if num > maxCount {
			maxCount = num
			minYear = year
		} else if num == maxCount && year < minYear {
			minYear = year
		}
	}

	return minYear
}
