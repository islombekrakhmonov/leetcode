package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(maxProduct(345))
}

func maxProduct(n int) int {
	var array []int
	nStr := strconv.Itoa(n)
	for i := 0; i < len(nStr); i++ {
		nInt, _ := strconv.Atoi(string(nStr[i]))
		array = append(array, nInt)
	}

	sort.Slice(array, func(i, j int) bool {
		return array[i] > array[j]
	})

	return array[1] * array[0]
}
