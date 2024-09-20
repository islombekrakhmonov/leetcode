package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(sumOfEncryptedInt([]int{23, 234, 543, 54}))
}

func sumOfEncryptedInt(nums []int) int {

	var sum int

	for _, v := range nums {
		num := encrypt(v)
		sum += num
	}

	return sum
}

func encrypt(i int) int {

	iStr := strconv.Itoa(i)
	var max int
	var outputStr string
	var output int

	for _, v := range iStr {
		integer, _ := strconv.Atoi(string(v))
		if max < integer {
			max = integer
		}
	}

	maxStr := strconv.Itoa(max)

	for range iStr {
		outputStr += maxStr
	}

	output, _ = strconv.Atoi(outputStr)

	return output
}

func numberToSlice(n int) int {
	var result int

	for n > 0 {
		digit := n % 10
		result += digit
		n = n / 10
	}

	return result
}
