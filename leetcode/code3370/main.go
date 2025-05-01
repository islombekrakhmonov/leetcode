package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(smallestNumber(5))
}

func smallestNumber(n int) int {

	for {
		binaryString := strconv.FormatInt(int64(n), 2)
		if strings.Contains(binaryString, "0") {
			n++
		} else {
			return n
		}
	}
}
