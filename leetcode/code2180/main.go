package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countEven(30))
}

func countEven(num int) int {
	var output int
	var sum int
    for i:=1; i<=num; i++{
		sum = 0
		numStr := strconv.Itoa(i)
		for _,v := range numStr {
			iStr,_ := strconv.Atoi(string(v))
			sum += iStr
		}

		if sum %2 ==0{
			output++
		}
	}

	return output
}