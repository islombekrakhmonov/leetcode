package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(maximumOddBinaryNumber("0101"))
}

func maximumOddBinaryNumber(s string) string {
	ones := strings.Count(s, "1")
	zeros := strings.Count(s, "0")
	result := strings.Repeat("1", ones-1) + strings.Repeat("0", zeros) + "1"


	return result
}
