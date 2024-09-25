package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(hammingWeight(2147483645))
}

func hammingWeight(n int) int {
	var output int

	binaryStr := fmt.Sprintf("%b", n)

	for i := 0; i < len(binaryStr); i++ {

		bit, _ := strconv.Atoi(string(binaryStr[i]))
		if bit == 1 {
			output++
		}
	}

	return output
}
