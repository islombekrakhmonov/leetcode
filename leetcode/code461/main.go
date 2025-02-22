package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(hammingDistance(1, 2))
}

func hammingDistance(x int, y int) int {

	var count int
	bit1, bit2 := equalizeBitLengths(x, y)

	for i := 0; i < len(bit1); i++ {
		if bit1[i] != bit2[i] {
			count++
		}
	}

	return count
}

func equalizeBitLengths(num1, num2 int) (string, string) {
	bin1 := strconv.FormatInt(int64(num1), 2)
	bin2 := strconv.FormatInt(int64(num2), 2)

	maxLen := len(bin1)
	if len(bin2) > maxLen {
		maxLen = len(bin2)
	}

	bin1 = strings.Repeat("0", maxLen-len(bin1)) + bin1
	bin2 = strings.Repeat("0", maxLen-len(bin2)) + bin2

	return bin1, bin2
}
