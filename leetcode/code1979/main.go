package main

import "fmt"

func main() {
	fmt.Println(findGCD([]int{2,5,6,9,10}))
}

func findGCD(nums []int) int {
	min,max := nums[0],1
	var output int
	for i := 0; i<len(nums); i++{
		if nums[i] < min {
			min = nums[i]
		} 
		if max < nums[i] {
			max = nums[i]
		}
	}
	for i:= max; i>0; i--{
		if min % i == 0 && max % i == 0 {
			output = i
			break
		}
	}
    return output
}