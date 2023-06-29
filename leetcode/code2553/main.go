package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(separateDigits([]int{13,25,83,77}))
}

func separateDigits(nums []int) []int {
		var output []int
		var value int
		var strNums []string 
		for _,v := range nums{
			strNums = append(strNums, strconv.Itoa(v))
		}
		for _,v := range strNums{
			for _,l := range v{
			value , _ = strconv.Atoi(string(l))
			output = append(output, value)
		}
	}
	return output 
}