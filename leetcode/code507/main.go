package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(checkPerfectNumber(28))
}

func checkPerfectNumber(num int) bool {
	if num <= 1 {
		return false
	}

	sqrt := int(math.Sqrt(float64(num)))
	sum := 1

	for i := 2; i <= sqrt; i++ {
		if num%i == 0 {
			fmt.Println(sum, " ", i)
			sum += i + num/i
		}
	}

	return sum == num
}
