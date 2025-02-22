package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(isBalanced("1234"))
}

func isBalanced(num string) bool {

	var sumOdd, sumEven int

	for i, v := range num {
		vInt, _ := strconv.Atoi(string(v))
		if i%2 == 0 {
			sumEven += vInt
		} else {
			sumOdd += vInt
		}
	}

	return sumEven == sumOdd
}
