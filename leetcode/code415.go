package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(addStrings("11", "111"))
}

func addStrings(num1 string, num2 string) string {
    num1Int, _ := strconv.Atoi(num1)
    num2Int, _ := strconv.Atoi(num2)
	sum := num1Int+num2Int
	return strconv.Itoa(sum)
}