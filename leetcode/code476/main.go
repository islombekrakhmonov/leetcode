package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(findComplement(5))
}

func findComplement(num int) int {
	var outputStr string

	bits := strconv.FormatInt(int64(num), 2)

	for _, bit := range bits {
		if bit == '0' {
			outputStr += string('1')
		} else {
			outputStr += string('0')
		}
	}

	output, _ := strconv.ParseInt(outputStr, 2, 64)

	return int(output)
}