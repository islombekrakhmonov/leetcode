package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isSameAfterReversals(609576))
}


func isSameAfterReversals(num int) bool {
	var first string
	var second string
	var foundNonZero bool
    number := strconv.Itoa(num)
	for i:=len(number)-1; i>=0; i--{
		if number[i] == '0' && !foundNonZero{
			continue
		} else {
			first += string(number[i])
			foundNonZero = true
		}
	}
	fmt.Println(first)

	for i:=len(first)-1; i>=0; i--{
		second += string(first[i])
	}
	fmt.Println(second)
	digitInt, _ := strconv.Atoi(second)
	fmt.Println(digitInt)
	if num == digitInt{
		return true
	}
	return false
}