package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(smallestNumber(15, 3))
}

func smallestNumber(n int, t int) int {

	for {
		nStr := strconv.Itoa(n)
		product := 1
		for _, v := range nStr {
			vInt, _ := strconv.Atoi(string(v))
			product *= vInt
		}
		if product%t == 0 {
			break
		}
		n++
	}

	return n
}
