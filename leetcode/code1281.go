package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(subtractProductAndSum(234))
}

func subtractProductAndSum(n int) int {
    var slc []int
	product := 1
	var sum int
	number := strconv.Itoa(n)
	for _, digit := range number {
		digitInt, _ := strconv.Atoi(string(digit))
		slc = append(slc, digitInt)
	}

	for _,v := range slc{
		product *= v
		sum += v
	}

	return product - sum
}