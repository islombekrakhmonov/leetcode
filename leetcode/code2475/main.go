package main

import "fmt"

func main() {
	fmt.Println(unequalTriplets([]int{4,4,2,4,3}))
}

func unequalTriplets(nums []int) int {
	var output int 
    for i:=0;i<len(nums);i++{
		for j:=i+1; j<len(nums);j++{
			for k:=j+1; k<len(nums);k++{
				if nums[i] != nums[j] && nums[i] != nums[k] && nums[j] != nums[k]{
					output++
				}
			}
		}
	}
	return output
}