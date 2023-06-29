package main

import "fmt"

func main() {
	fmt.Println(maximizeSum([]int{4,4,9,10,10,9,3,8,4,2,5,3,8,6,1,10,4,5,3,2,3,9,5,7,10,4,9,10,1,10,4}, 6))
}

func maximizeSum(nums []int, k int) int {
	var output int 
	var max int 
	for _,v := range nums{
		if v > max {
			max = v
		}
	}
    for i:=0;i<k;i++{
		output += max
		max +=1
	}
	return output
}