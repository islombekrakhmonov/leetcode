package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumSum(2932))
}

func minimumSum(num int) int {
	var slice []int

	for num > 0{
		slice = append(slice, num %10)
		num /= 10
	}

	sort.Ints(slice)
	new1 := slice[0] * 10 + slice[2]
	new2 := slice[1] * 10 + slice[3]

	return new1+new2
}