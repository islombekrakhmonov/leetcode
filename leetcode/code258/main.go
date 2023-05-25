package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(addDigits(38))	
}



func addDigits(num int) int {
    number := strconv.Itoa(num)
	for len(number) > 1 {
		num = 0
		for _,v := range number {
			digit, _ := strconv.Atoi(string(v))
			num += digit
		}
		number = strconv.Itoa(num)
	}
	return num 
}
