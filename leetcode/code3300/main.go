package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(minElement([]int{10, 12, 13, 14}))
}

func minElement(nums []int) int {

	var min = math.MaxInt
	sumOfDigitsSlice := make([]int, len(nums))

	for i, num := range nums {
		sum := sumOfDigits(num)
		sumOfDigitsSlice[i] = sum
		if min > sum {
			min = sum
		}
	}

	return min
}

func sumOfDigits(i int) (sum int) {
	iStr := strconv.Itoa(i)

	for _, v := range iStr {
		num, _ := strconv.Atoi(string(v))
		sum += num
	}

	return sum
}
