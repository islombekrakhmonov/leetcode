package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(evenOddBit(2))
}

func evenOddBit(n int) []int {
	binaryStr := strconv.FormatInt(int64(n), 2)

	fmt.Println(0%2==0)
	fmt.Println(binaryStr)
	var output [2]int
	var even, odd = 0,0
	for i,v := range binaryStr {
		if v == '1' {
			if i % 2 == 0 {
				odd++
			} else {
				even++
			}
		}
	}

	output[0] = odd
	output[1] = even

	return output[:]
}