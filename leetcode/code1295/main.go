package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(findNumbers([]int{12,345,2,6,7896}))
}

func findNumbers(nums []int) int {
	var count int
	var output int
	for _,v := range nums{
		count = 0
		str := strconv.Itoa(v)
		for i:=0; i<len(str); i++{
			count++
		}
		if count % 2 ==0 {
			output++
		}
	}
    
	return output
}