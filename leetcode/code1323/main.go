package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(maximum69Number(9669))
}

func maximum69Number (num int) int {
	max := num
	number := strconv.Itoa(num)
	for i:=0; i<len(number); i++{
		if number[i] == '6'{
			number = number[:i] + "9" + number[i+1:]
			digitInt, _ := strconv.Atoi(number)
			if digitInt > max{
				max = digitInt
			}
			number = strconv.Itoa(num)
		} else {
			number = number[:i] + "6" + number[i+1:]
			digitInt, _ := strconv.Atoi(number)
			if digitInt > max{
				max = digitInt
			}
			number = strconv.Itoa(num)
		} 
	}
	return max
}