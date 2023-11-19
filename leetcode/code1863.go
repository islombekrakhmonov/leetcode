package main

import "fmt"

func main() {
	fmt.Println(subsetXORSum([]int{5,1,6}))
}

func subsetXORSum(nums []int) int {
	var output int
	for subsetMask := 0; subsetMask < (1 << len(nums)); subsetMask++ {
        result := 0
        for i, num := range nums {
            if (subsetMask & (1 << i)) != 0 {
                result ^= num
            }
        }
        output += result
    }
	return output
}