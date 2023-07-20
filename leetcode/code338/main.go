package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countBits(2))
}

func countBits(n int) []int {
	var output []int
	for i:=0; i<n+1;i++{
		count := 0
		binary := strconv.FormatInt(int64(i), 2)
		for _,k:= range binary{
			if string(k) == "1"{ 
				count++
			}
		}
		output = append(output, count)
	}
    return output
}