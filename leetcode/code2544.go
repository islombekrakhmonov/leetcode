package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(alternateDigitSum(886996))
	fmt.Println(alternateDigitSum1(886996))
}

func alternateDigitSum(n int) int {
	nStr := strconv.Itoa(n)
	first, _ := strconv.Atoi(string(nStr[0]))
	for i:=1; i<len(nStr); i++{
		if i % 2 == 0 {
			num, _ := strconv.Atoi(string(nStr[i]))
			first += num
		} else {
			num, _ := strconv.Atoi(string(nStr[i]))
			first -= num
		}
	}

	return first
}

func alternateDigitSum1(n int) int {
	var result int
	for ; n != 0; n /= 10 {
		fmt.Println("N", n)
		result = n%10 - result
		fmt.Println("result", result)
	}
	return result
}