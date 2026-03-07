package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(decimalRepresentation(531))
}

func decimalRepresentation(n int) []int {
	var output []int

	nStr := fmt.Sprintf("%v", n)

	for i := len(nStr); i > 0; i-- {
		digit := n % 10
		if digit != 0 {
			output = append(output, digit*int(math.Pow10(len(nStr)-i)))
		}
		n /= 10
	}

	sort.Slice(output, func(i, j int) bool {
		return output[i] > output[j]
	})

	return output
}

//i think the problem description can be more clear . we simple need to abstract the digit and if the digit is not zero then we need to push digit*pow(10,digit_position_from_right) to the answer , at the end we will reverse the answer array.
