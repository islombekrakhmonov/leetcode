package main

import "fmt"

func main() {
	fmt.Println(xorOperation(5,0))
}

func xorOperation(n int, start int) int {
	var nums []int
	var output int
	for i:=0; i<n;i++{
		nums = append(nums, start + 2 * i)
		output ^= nums[i]
	}
    
	return output
}