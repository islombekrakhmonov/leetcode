package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(minimumDistance([]int{1, 1, 2, 3, 2, 1, 2}))
}

func minimumDistance(nums []int) int {
	occurenceMap := make(map[int][]int)

	for i := 0; i < len(nums); i++ {
		array := occurenceMap[nums[i]]
		array = append(array, i)
		occurenceMap[nums[i]] = array
	}

	var min = math.MaxInt

	for _, array := range occurenceMap {
		if len(array) < 3 {
			continue
		}

		for i := 0; i < len(array)-2; i++ {
			distance := math.Abs(float64(array[i]-array[i+1])) + math.Abs(float64(array[i+1]-array[i+2])) + math.Abs(float64(array[i+2]-array[i]))
			if distance < float64(min) {
				min = int(distance)
			}
		}
	}

	if min == math.MaxInt {
		return -1
	}

	return min
}
