package main

import "fmt"

func main() {

	fmt.Println(twoSum([]int{2,5,5,11}, 10))
}


func twoSum(nums []int, target int) []int {
	var output []int
	for k,v := range nums{
		for i:=k+1;i<len(nums); i++{
			if v+nums[i] == target{
				output = append(output, k)
				output = append(output, i)
				return output
			}
		}
	}
	 
	return output
}