package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(sumIndicesWithKSetBits([]int{5,10,1,5,2}, 1))
}

func sumIndicesWithKSetBits(nums []int, k int) int {
	var output int
    for i:=0; i<len(nums); i++{
		iBinary := strconv.FormatInt(int64(i), 2)
		count := 0
		for _,v := range iBinary {
			if v == '1' {
				count ++
			}
		}
		if count == k {
			output += nums[i]
		}
	}
	return output
}