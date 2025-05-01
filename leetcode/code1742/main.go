package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(countBalls(52, 61))
}

func countBalls(lowLimit int, highLimit int) int {

	frequency := make(map[int]int)
	for i := lowLimit; i <= highLimit; i++ {
		frequency[sumOfDigits(i)]++
	}

	max := 0

	for _, v := range frequency {
		if v > max {
			max = v
		}
	}

	return max
}

func sumOfDigits(num int) (sum int) {
	numStr := strconv.Itoa(num)

	for _, value := range numStr {
		valueInt, _ := strconv.Atoi(string(value))
		sum += valueInt
	}

	return sum
}
