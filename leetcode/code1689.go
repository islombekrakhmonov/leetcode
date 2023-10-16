package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(minPartitions("82734"))
}

func minPartitions(n string) int {
    num, _ := strconv.Atoi(n)

	var sum int
	tens := num%10
	sum = tens * 11
	elevens := ((num-sum)/10)
	sum = elevens * 10 + sum

	maxDigit := 0
	for i := 0; i < len(n); i++ {
        digit := int(n[i] - '0')
		if digit > maxDigit {
            maxDigit = digit
        }
	}

	fmt.Println(maxDigit)
	return tens+elevens
}