package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(generateKey(1, 10, 1000))
}

func generateKey(num1 int, num2 int, num3 int) int {

	var output int
	var outputStr string
	num1Str := fmt.Sprintf("%04d", num1)
	num2Str := fmt.Sprintf("%04d", num2)
	num3Str := fmt.Sprintf("%04d", num3)

	for i := 0; i < 4; i++ {
		minNum := min(num1Str[i], num2Str[i], num3Str[i])
		outputStr += string(minNum)
	}

	output, _ = strconv.Atoi(outputStr)

	return output
}

func min(num1, num2, num3 byte) byte {
	if num1 <= num2 && num1 <= num3 {
		return num1
	} else if num2 <= num3 {
		return num2
	}
	return num3
}
