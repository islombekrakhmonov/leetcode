package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findClosestNumber([]int{-4, -2, 1, 4, 8, -1}))

}

func findClosestNumber(nums []int) int {

	numDistance := make(map[int]int)
	min := math.MaxInt
	minDistanceArray := []int{}
	greaterNum := math.MinInt

	for _, num := range nums {
		distance := int(math.Abs(float64(num) - 0))
		if distance < min {
			min = distance
		}
		numDistance[num] = distance
	}

	for num, distance := range numDistance {
		if distance == min {
			minDistanceArray = append(minDistanceArray, num)
		}
	}

	for _, num := range minDistanceArray {
		if num > greaterNum {
			greaterNum = num
		}

	}

	return greaterNum
}
