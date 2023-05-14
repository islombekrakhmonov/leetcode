package main

import "fmt"

func main() {
	number := []int{2,4,2}
	fmt.Println(singleNumber(number))
	
}

func singleNumber(nums []int) int {
	res := nums[0]

	if len(nums) == 1 {
		return res
	}else {
		for i := 1; i < len(nums); i++ {
			res ^= nums[i]
		}
	}
return res
}