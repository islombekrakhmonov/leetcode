package main

import "fmt"

func main() {
	fmt.Println(shuffle([]int{2,5,1,3,4,7}, 3))
}

func shuffle(nums []int, n int) []int {
    point := len(nums) / 2
	var x []int
	var y []int
	var output []int
	for i:=0; i<point; i++{
		x = append(x, nums[i])
	}
	for i:=point; i<len(nums); i++{
		y = append(y, nums[i])
	}
	for i:=0; i<point; i++{
		output = append(output, x[i])
		output = append(output, y[i])
	}
	return output
}