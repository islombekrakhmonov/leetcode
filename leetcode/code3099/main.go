package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(sumOfTheDigitsOfHarshadNumber(18))
}

func sumOfTheDigitsOfHarshadNumber(x int) int {

	xStr := strconv.Itoa(x)
	var sum int

	for _, v := range xStr {
		num, _ := strconv.Atoi(string(v))
		sum += num
	}

	if x%sum == 0 {
		return sum
	}

	return -1
}
