package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(getLucky("zbax", 2))
}

func getLucky(s string, k int) int {
	var output int

	alphabet := "abcdefghijklmnopqrstuvwxyz"

	alphaberMap := make(map[rune]string)

	for i, v := range alphabet {
		iStr := strconv.Itoa(i + 1)
		alphaberMap[v] = iStr
	}

	var numStr string
	for _, v := range s {
		numStr += alphaberMap[v]
	}

	for k > 0 {
		var sum int
		for _, num := range numStr {
			numInt, _ := strconv.Atoi(string(num))
			sum += numInt
		}
		output = sum

		numStr = strconv.Itoa(output)
		k--
	}

	return output
}
