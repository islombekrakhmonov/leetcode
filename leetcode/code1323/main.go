package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(maximum69Number(9669))
}

func maximum69Number(num int) int {
	max := num

	numStr := strconv.Itoa(num)

	for i := 0; i < len(numStr); i++ {
		if numStr[i] == '6' {
			formattedNum := numStr[:i] + "9" + numStr[i+1:]
			formattedNumInt, _ := strconv.Atoi(formattedNum)
			if formattedNumInt > max {
				max = formattedNumInt
			}
		}
	}

	return max
}
