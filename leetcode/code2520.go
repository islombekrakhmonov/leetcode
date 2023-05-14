package main

import (
	"fmt"
)

func main() {
	input := 121
	fmt.Println(countDigits(input))
}

func countDigits(num int) int {
	var output int
	s := num
	slc := []int{}
	for num > 0 {
		slc = append(slc, num%10)
		num /= 10
	}
	for i:=0; i<len(slc); i++{
		if s % slc[i] == 0{
			output ++
		}
	}
	return output
}
