package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(average([]int{4000, 3000, 2000, 1000}))
	// fmt.Println(average([]int{1000, 2000, 3000}))

}

func average(salary []int) float64 {

	min := math.MaxInt
	max := math.MinInt
	sum := 0

	for _, s := range salary {
		if s < min {
			min = s
		}
		if s > max {
			max = s
		}
		sum += s
	}

	sum = sum - max - min

	return float64(sum) / float64(len(salary)-2)
}
