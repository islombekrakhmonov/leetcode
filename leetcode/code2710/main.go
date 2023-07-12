package main

import (
	"fmt"
	"math"
)

func removeTrailingZeros(num string) string {
	if len(num) == 0 {
		return num
	}

	for string(num[len(num)-1]) == "0" {
		num = num[:len(num)-1]
		fmt.Println(num)
	}

	return num
}

func main() {
	num := "10000"
	result := removeTrailingZeros(num)
	fmt.Println("Result:", result)
	fmt.Println("Maximum value of int32:", math.MaxInt32)

}