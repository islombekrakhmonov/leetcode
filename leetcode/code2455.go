package main

import "fmt"

func main() {
	fmt.Println(averageValue([]int{1,3,6,10,12,15}))
}

func averageValue(nums []int) int {
	sum, count := 0, 0
	for i:=0; i<len(nums); i++{
		if nums[i] % 2 == 0 && nums[i] % 3 == 0 {
			count++
			sum += nums[i]
		}
	}
    
	return sum/count
}