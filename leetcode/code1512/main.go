package main

import "fmt"

func main() {
	fmt.Println(numIdenticalPairs([]int{1,2,3,1,1,3}))
}

func numIdenticalPairs(nums []int) (output int) {
    for i:=0;i<len(nums);i++{
		for j:=i+1;j<len(nums);j++{
			if nums[i] == nums[j] && i<j{
				output++
			}
		}
	}
	return 
}